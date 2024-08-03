package controller

import (
	"context"
	"fmt"
	"net/http"

	"forum/internal/models"
	"forum/pkg/logger"
)

type userService interface {
	Register(ctx context.Context, user models.User) error
}

type userController struct {
	service userService
	log     *logger.Logger
}

func newUserController(userService userService, log *logger.Logger) *userController {
	return &userController{
		service: userService,
		log:     log,
	}
}

func (c *userController) RegisterGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a HTML form for signing up a new user...")
}

func (c *userController) RegisterPOST(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO: implement error handling/rendering
		return
	}

	if err := r.ParseForm(); err != nil {
		c.log.Error(fmt.Sprintf("failed to parse form: %v", err.Error()))
		return
	}

	email := r.Form.Get("email")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	passwordConfirm := r.Form.Get("password_confirm")

	if password != passwordConfirm {
		c.log.Warn("passwords do not match")
		return
	}

	if email == "" {
		c.log.Warn("email is empty")
		return
	}

	if username == "" {
		c.log.Warn("username is empty")
		return
	}

	if err := c.service.Register(r.Context(), models.User{
		Email:    email,
		Username: username,
		Password: password,
	}); err != nil {
		c.log.Error(fmt.Sprintf("failed to register user: %v", err.Error()))
		return
	}
}

func (c *userController) LoginGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a HTML form for logging in a user...")
}

func (c *userController) LoginPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (c *userController) LogoutPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
