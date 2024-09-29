package repository

import (
	"database/sql"

	"github.com/anne-markis/fermtrack/internal/domain"
)

type MySQLFermentationRepository struct {
	db *sql.DB
}

func NewMySQLFermentationRepository(db *sql.DB) *MySQLFermentationRepository {
	return &MySQLFermentationRepository{db}
}

func (r *MySQLFermentationRepository) FindAll() ([]domain.Fermentation, error) {
	rows, err := r.db.Query("SELECT id, uuid, nickname, start_at, bottled_at, recipe_notes, tasting_notes, deleted_at FROM fermentations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fermentations []domain.Fermentation
	for rows.Next() {
		var fermentation domain.Fermentation
		if err := rows.Scan(&fermentation.ID, &fermentation.UUID, &fermentation.Nickname, &fermentation.StartAt, &fermentation.BottledAt, &fermentation.RecipeNotes, &fermentation.TastingNotes, &fermentation.DeletedAt); err != nil {
			return nil, err
		}
		fermentations = append(fermentations, fermentation)
	}
	return fermentations, nil
}

func (r *MySQLFermentationRepository) FindByID(uuid string) (*domain.Fermentation, error) {
	// Implement MySQL query for fetching by UUID
	return nil, nil
}

func (r *MySQLFermentationRepository) Create(f *domain.Fermentation) error {
	// Implement MySQL insert query
	return nil
}

func (r *MySQLFermentationRepository) Update(f *domain.Fermentation) error {
	// Implement MySQL update query
	return nil
}

func (r *MySQLFermentationRepository) Delete(uuid string) error {
	// Implement MySQL delete query
	return nil
}
