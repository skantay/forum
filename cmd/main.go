package main

import "net/http"

func main() {
	// TODO: configure server
	server := http.Server{
		//TODO: take config envs from .env
		Addr: ":8080",
	}

	// TODO: implement all layers

	// TODO: implement gracefull shutdown
	server.ListenAndServe()
}
