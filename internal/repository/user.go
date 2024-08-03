package repository

import (
	"context"
	"database/sql"
	"fmt"
	"forum/pkg/logger"
)

type UserRepository struct {
	db  *sql.DB
	log *logger.Logger
}

func NewUserRepository(db *sql.DB, log *logger.Logger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (r *UserRepository) Register(ctx context.Context, email string, username string, password []byte) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(email, username, password) VALUES($1, $2, $3)", email, username, password)
	if err != nil {
		r.log.Error("failed to insert user")
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (r *UserRepository) Authenticate(email, password string) (int, error) {
    return 0, nil
}

// We'll use the Exists method to check if a user exists with a specific ID.
func (r *UserRepository) Exists(id int) (bool, error) {
    return false, nil
}