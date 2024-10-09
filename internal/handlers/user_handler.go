package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anne-markis/fermtrack/internal/app/domain"
	"github.com/anne-markis/fermtrack/internal/utils"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userRepo domain.UserRepository
}

func NewUserHandler(repo domain.UserRepository) *UserHandler {
	return &UserHandler{userRepo: repo}
}

// GetUser retrieves a user by the uuid passed in on the request
func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fermentation, err := u.userRepo.FindByUUID(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(fermentation)
}

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser creates a new user
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUserReq NewUserRequest
	if err := json.NewDecoder(r.Body).Decode(&newUserReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPwd, err := utils.HashPassword(newUserReq.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.userRepo.Create(newUserReq.Username, hashedPwd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser, err := u.userRepo.FindByUsername(newUserReq.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(newUser)
}
