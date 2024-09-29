package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// TODO this naming bothers me
type FermtrackServer struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *FermtrackServer {
	return &FermtrackServer{db: db}
}

type Fermentation struct {
	ID           int        `json:"id"`
	UUID         string     `json:"uuid"`
	Nickname     string     `json:"nickname"`
	StartAt      time.Time  `json:"start_at"`
	BottledAt    *time.Time `json:"bottled_at"`
	RecipeNotes  string     `json:"recipe_notes"`
	TastingNotes *string    `json:"tasting_notes"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type FermentationsResponse struct {
	Fermentations []Fermentation `json:"fermentations"`
	Page          int            `json:"page"`
	Limit         int            `json:"limit"`
	Total         int            `json:"total"`
}

func (fs *FermtrackServer) GetProjectHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uuid := vars["uuid"]

	switch req.Method {
	case http.MethodGet:
		fs.handleGetFermentation(w, req, uuid)
	case http.MethodPut:
		fs.handleUpdateFermentation(w, req, uuid)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ListProjectsHandler handles GET requests for fetching the list of fermentations
func (fs *FermtrackServer) ListProjectsHandler(w http.ResponseWriter, req *http.Request) { // TODO rename
	queryParams := req.URL.Query()

	// Parse optional filters
	startAt := queryParams.Get("start_at")
	bottledAt := queryParams.Get("bottled_at")
	deletedAt := queryParams.Get("deleted_at")

	// Pagination parameters (default page: 1, limit: 10)
	page := 1
	limit := 10
	if p, err := strconv.Atoi(queryParams.Get("page")); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(queryParams.Get("limit")); err == nil && l > 0 {
		limit = l
	}

	// Build SQL query TODO
	query := "SELECT id, uuid, nickname, start_at, bottled_at, recipe_notes, tasting_notes, deleted_at FROM fermentations WHERE 1=1"
	args := []interface{}{}

	// Add filtering clauses based on query parameters
	if startAt != "" {
		query += " AND start_at >= ?"
		args = append(args, startAt)
	}
	if bottledAt != "" {
		query += " AND bottled_at <= ?"
		args = append(args, bottledAt)
	}
	if deletedAt != "" {
		query += " AND deleted_at IS NOT NULL"
	}

	// Pagination
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	// Execute the query
	rows, err := fs.db.Query(query, args...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching fermentations: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process the result set
	var fermentations []Fermentation
	for rows.Next() {
		var fermentation Fermentation
		err := rows.Scan(&fermentation.ID, &fermentation.UUID, &fermentation.Nickname, &fermentation.StartAt, &fermentation.BottledAt, &fermentation.RecipeNotes, &fermentation.TastingNotes, &fermentation.DeletedAt)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning rows: %v", err), http.StatusInternalServerError)
			return
		}
		fermentations = append(fermentations, fermentation)
	}

	// Count total records for pagination
	var total int
	err = fs.db.QueryRow("SELECT COUNT(*) FROM fermentations WHERE 1=1").Scan(&total)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error counting fermentations: %v", err), http.StatusInternalServerError)
		return
	}

	// Construct the response
	response := FermentationsResponse{
		Fermentations: fermentations,
		Page:          page,
		Limit:         limit,
		Total:         total,
	}

	// Write response in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}

// handleGetFermentation retrieves a fermentation by UUID
func (ftServer *FermtrackServer) handleGetFermentation(w http.ResponseWriter, r *http.Request, uuid string) {
	fermentation := Fermentation{}
	query := "SELECT id, uuid, nickname, start_at, bottled_at, recipe_notes, tasting_notes, deleted_at FROM fermentations WHERE uuid = ?"
	err := ftServer.db.QueryRow(query, uuid).Scan(&fermentation.ID, &fermentation.UUID, &fermentation.Nickname, &fermentation.StartAt, &fermentation.BottledAt, &fermentation.RecipeNotes, &fermentation.TastingNotes, &fermentation.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Fermentation not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error retrieving fermentation: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fermentation); err != nil {
		http.Error(w, "Error encoding response as JSON", http.StatusInternalServerError)
	}
}

// handleUpdateFermentation updates the details of a fermentation by UUID
func (ftServer *FermtrackServer) handleUpdateFermentation(w http.ResponseWriter, r *http.Request, uuid string) {
	body, err := ioutil.ReadAll(r.Body) // TODO
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var updatedFermentation Fermentation
	if err := json.Unmarshal(body, &updatedFermentation); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	var existingFermentation Fermentation
	query := "SELECT id FROM fermentations WHERE uuid = ?"
	err = ftServer.db.QueryRow(query, uuid).Scan(&existingFermentation.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Fermentation not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error checking fermentation existence: %v", err), http.StatusInternalServerError)
		}
		return
	}

	updateQuery := "UPDATE fermentations SET "
	updateFields := []string{}
	args := []interface{}{}

	if updatedFermentation.Nickname != "" {
		updateFields = append(updateFields, "nickname = ?")
		args = append(args, updatedFermentation.Nickname)
	}
	if updatedFermentation.BottledAt != nil {
		updateFields = append(updateFields, "bottled_at = ?")
		args = append(args, updatedFermentation.BottledAt)
	}
	if updatedFermentation.RecipeNotes != "" {
		updateFields = append(updateFields, "recipe_notes = ?")
		args = append(args, updatedFermentation.RecipeNotes)
	}
	if updatedFermentation.TastingNotes != nil {
		updateFields = append(updateFields, "tasting_notes = ?")
		args = append(args, updatedFermentation.TastingNotes)
	}

	if updatedFermentation.DeletedAt != nil {
		updateFields = append(updateFields, "deleted_at = ?")
		args = append(args, *updatedFermentation.DeletedAt)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	updateQuery += fmt.Sprintf("%s WHERE uuid = ?", fmt.Sprintf("%s", updateFields))
	args = append(args, uuid)

	// Execute the update query
	_, err = ftServer.db.Exec(updateQuery, args...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating fermentation: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
