//go:generate mockery --name=FermentationRepository --dir=internal/app/domain --output=internal/app/mocks --with-expecter
package domain

import (
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"
)

type MySQLFermentationRepository struct {
	db *sql.DB
}

func NewMySQLFermentationRepository(db *sql.DB) *MySQLFermentationRepository {
	return &MySQLFermentationRepository{db}
}

func (r *MySQLFermentationRepository) FindAll() ([]Fermentation, error) {
	log.Info().Msg("fermentations - find all")
	rows, err := r.db.Query(`
	SELECT
		id
		, uuid
		, nickname
		, start_at
		, bottled_at
		, recipe_notes
		, tasting_notes
	FROM fermentations
		WHERE deleted_at IS NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fermentations []Fermentation
	for rows.Next() {
		var fermentation Fermentation
		if err := rows.Scan(
			&fermentation.ID,
			&fermentation.UUID,
			&fermentation.Nickname,
			&fermentation.StartAt,
			&fermentation.BottledAt,
			&fermentation.RecipeNotes,
			&fermentation.TastingNotes); err != nil {
			return nil, err
		}
		fermentations = append(fermentations, fermentation)
	}
	return fermentations, nil
}

func (r *MySQLFermentationRepository) FindByUUID(uuid string) (*Fermentation, error) {
	log.Info().Any("uuid", uuid).Msg("fermentations - find by uuid")
	query := `
	SELECT
		id
		, uuid
		, nickname
		, start_at
		, bottled_at
		, recipe_notes
		, tasting_notes
	FROM fermentations
		WHERE uuid = ?
	`

	var fermentation Fermentation
	if err := r.db.QueryRow(query, uuid).Scan(
		&fermentation.ID,
		&fermentation.UUID,
		&fermentation.Nickname,
		&fermentation.StartAt,
		&fermentation.BottledAt,
		&fermentation.RecipeNotes,
		&fermentation.TastingNotes,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}

	return &fermentation, nil
}
