//go:generate mockery --name=FermentationRepository --dir=internal/repository --output=internal/mocks --with-expecter
package repository

import (
	"database/sql"
)

type MySQLFermentationRepository struct {
	db *sql.DB
}

func NewMySQLFermentationRepository(db *sql.DB) *MySQLFermentationRepository {
	return &MySQLFermentationRepository{db}
}

func (r *MySQLFermentationRepository) FindAll() ([]Fermentation, error) {
	rows, err := r.db.Query("SELECT id, uuid, nickname, start_at, bottled_at, recipe_notes, tasting_notes, deleted_at FROM fermentations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fermentations []Fermentation
	for rows.Next() {
		var fermentation Fermentation
		if err := rows.Scan(&fermentation.ID, &fermentation.UUID, &fermentation.Nickname, &fermentation.StartAt, &fermentation.BottledAt, &fermentation.RecipeNotes, &fermentation.TastingNotes, &fermentation.DeletedAt); err != nil {
			return nil, err
		}
		fermentations = append(fermentations, fermentation)
	}
	return fermentations, nil
}

func (r *MySQLFermentationRepository) FindByID(uuid string) (*Fermentation, error) {
	// Implement MySQL query for fetching by UUID
	return nil, nil
}

func (r *MySQLFermentationRepository) Create(f *Fermentation) error {
	// Implement MySQL insert query
	return nil
}

func (r *MySQLFermentationRepository) Update(f *Fermentation) error {
	// Implement MySQL update query
	return nil
}

func (r *MySQLFermentationRepository) Delete(uuid string) error {
	// Implement MySQL delete query
	return nil
}
