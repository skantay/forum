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

func (r *UserRepository) Register(ctx context.Context, email string, password []byte) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(email, password) VALUES($1, $2)", email, password)
	if err != nil {
		r.log.Error("failed to insert user")
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
