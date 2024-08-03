package controller

import (
	"net/http"

	"forum/pkg/logger"
)

func NewAPI(
	userService userService,
	log *logger.Logger,
) *http.ServeMux {
	userController := newUserController(userService, log)

	mux := http.NewServeMux()

	mux.HandleFunc("/register", userController.RegisterHandler)
	mux.HandleFunc("/login", userController.LoginHandler)
	mux.HandleFunc("/logout", userController.LogoutPOST)

	return mux
}

func (uc *userController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uc.RegisterGET(w, r)
	case http.MethodPost:
		uc.RegisterPOST(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (uc *userController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uc.LoginGET(w, r)
	case http.MethodPost:
		uc.LoginPOST(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
