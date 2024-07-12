package controller

import (
	"forum/internal/models"
	"forum/internal/render"
	"net/http"
)

type userService interface {
	Register(registration models.UserRegistration) error
}

type User struct {
	r render.Render
	u userService
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {
}
