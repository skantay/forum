package controller

import (
	"forum/internal/render"
	"net/http"
)

type User struct {
	r render.Render
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) error {
}
