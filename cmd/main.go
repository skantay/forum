package main

import (
	"fmt"
	"forum/internal/controller"
	"net/http"
	"time"
)

func main() {
	handler := controller.New()

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
