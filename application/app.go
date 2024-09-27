package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/hanif-adedotun/ecommerce-golang/db"
)

type App struct {
	router http.Handler
	db     *sql.DB
}

func New() *App {
	dbConn, err := db.ConnectDB()
	if err != nil {
		// Handle the error (e.g., log it and exit)
		log.Fatal(err)
	}
	app := &App{
		db: dbConn,
	}

	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	defer func() {
		if err := a.db.Close(); err != nil {
			fmt.Println("Failed to close Postgres", err)
		}
	}()

	fmt.Println("Starting server...")

	ch := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("could not start server: %w", err)
		}
		close(ch)
	}()
	var err error // Declare err variable
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
