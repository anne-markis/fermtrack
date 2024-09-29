package domain

import "time"

type Fermentation struct {
	ID           int
	UUID         string
	Nickname     string
	StartAt      time.Time
	BottledAt    *time.Time
	RecipeNotes  string
	TastingNotes *string
	DeletedAt    *time.Time
}

type FermentationRepository interface {
	FindAll() ([]Fermentation, error)
	FindByID(uuid string) (*Fermentation, error)
	Create(fermentation *Fermentation) error
	Update(fermentation *Fermentation) error
	Delete(uuid string) error
}
