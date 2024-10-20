package domain

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db}
}

func (r *MySQLUserRepository) Create(username string, hashedPass string) error {
	log.Info().Any("username", username).Msg("users - create")

	query := `
	INSERT INTO users 
		(uuid, username, password, created_at)
	VALUES
		(?, ?, ?, now())`
	_, err := r.db.Exec(query, uuid.NewString(), username, hashedPass)
	return err
}

func (r *MySQLUserRepository) FindByUsername(username string) (*User, error) {
	log.Info().Any("username", username).Msg("users - find by username")
	query := `
	SELECT
		id
		, uuid
		, username
		, password
		, created_at
		, updated_at
		, deleted_at
	FROM users
		WHERE username = ?
	`

	var user User
	if err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.UUID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}
	return &user, nil
}

func (r *MySQLUserRepository) FindByUUID(uuid string) (*User, error) {
	log.Info().Any("uuid", uuid).Msg("users - find by uuid")
	query := `
	SELECT
		id
		, uuid
		, username
		, password
		, created_at
		, updated_at
		, deleted_at
	FROM users
		WHERE uuid = ?
	`

	var user User
	if err := r.db.QueryRow(query, uuid).Scan(
		&user.ID,
		&user.UUID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}
	return &user, nil
}
