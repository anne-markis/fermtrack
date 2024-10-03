package domain

import "time"

type Fermentation struct {
	ID           int        `json:"id"`
	UUID         string     `json:"uuid"`
	Nickname     string     `json:"nickname"`
	StartAt      time.Time  `json:"start_at"`
	BottledAt    *time.Time `json:"bottled_at"`
	RecipeNotes  string     `json:"recipe_notes"`
	TastingNotes *string    `json:"tasting_notes"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type FermentationRepository interface {
	FindAll() ([]Fermentation, error)
	FindByUUID(uuid string) (*Fermentation, error)
}
