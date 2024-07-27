package controller

import (
	"context"
	"fmt"
	"forum/internal/models"
	"forum/pkg/logger"
	"net/http"
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

	if err := c.service.Register(r.Context(), models.User{
		Email:    email,
		Password: password,
	}); err != nil {
		c.log.Error(fmt.Sprintf("failed to register user: %v", err.Error()))
		return
	}
}
