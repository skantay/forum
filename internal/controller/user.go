package controller

import (
	"forum/internal/models"
	"forum/internal/render"
	"net/http"
)

type userService interface {
	Register(registration models.UserRegistration) error
}

type user struct {
	r render.Render
	u userService
}

func (u *user) register(w http.ResponseWriter, r *http.Request) {
}
