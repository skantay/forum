package main

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/controller"
	"forum/internal/repository"
	"forum/internal/service"
	"forum/pkg/logger"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log := logger.NewLogger()
	db, err := initDB("forum.db")
	if err != nil {
		fmt.Println("failed to db", err.Error())
		log.Fatal(fmt.Sprintf("failed to init db: %v", err))

		return
	}

	userRepository := repository.NewUserRepository(db, log)

	userService := service.NewUserService(userRepository, log)

	mux := controller.NewAPI(userService, log)

	// TODO: configure server
	server := http.Server{
		//TODO: take config envs from .env
		Addr:    ":8080",
		Handler: mux,
	}

	// TODO: implement all layers

	// TODO: implement gracefull shutdown
	server.ListenAndServe()
}

func initDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	stmts, err := os.ReadFile("./migrations/up.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = db.ExecContext(ctx, string(stmts))
	if err != nil {
		return nil, fmt.Errorf("failed to execute migration: %w", err)
	}

	return db, nil
}
