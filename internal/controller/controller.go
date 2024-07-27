package controller

import (
	"forum/pkg/logger"
	"net/http"
)

func NewAPI(
	userService userService,
	log *logger.Logger,
) *http.ServeMux {
	userController := newUserController(userService, log)

	mux := http.NewServeMux()

	mux.HandleFunc("/register", userController.RegisterPOST)

	return mux
}
