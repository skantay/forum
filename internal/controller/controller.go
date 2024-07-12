package controller

import "net/http"

type controller struct {
	u user
}

func New() *http.ServeMux {
	mux := http.NewServeMux()

	ctrl := controller{}

	mux.HandleFunc("sign_up", ctrl.u.register)

	return mux
}
