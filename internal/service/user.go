package service

import (
	"context"
	"fmt"
	"forum/internal/models"
	"forum/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Register(ctx context.Context, email string, password []byte) error
}

type userService struct {
	repo UserRepository
	log  *logger.Logger
}

func NewUserService(userRepository UserRepository, log *logger.Logger) *userService {
	return &userService{
		repo: userRepository,
		log:  log,
	}
}

func (s *userService) Register(ctx context.Context, user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("failed to generate hash from password")
		return fmt.Errorf("failed to generate hash from password: %w", err)
	}

	if err := s.repo.Register(ctx, user.Email, hashed); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}
