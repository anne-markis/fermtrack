package domain

import "time"

type User struct {
	ID        int        `json:"id"`
	UUID      string     `json:"uuid"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserRepository interface {
	FindByUsername(username string) (*User, error)
	FindByUUID(uuuid string) (*User, error)
	Create(username string, password string) error
}
