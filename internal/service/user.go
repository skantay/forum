package service

import "forum/internal/models"

type userRepository interface {
	SaveUser(registration models.UserRegistration) error
}

type User struct {
	u userRepository
}

func (u *User) Register(user models.UserRegistration) error {
	return nil
}
