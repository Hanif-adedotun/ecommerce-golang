package application

import (
	"context"
	"fmt"
	"net/http"
	"database/sql"
"log"
	db "github.com/hanif-adedotun/ecommerce-golang/db"
)


type App struct {
	router http.Handler
	db *sql.DB
}

func New() *App {
	dbConn, err := db.ConnectDB();
	if err != nil {
		// Handle the error (e.g., log it and exit)
		log.Fatal(err)
	}
	app := &App{
		router: loadRoutes(),
		db: dbConn,
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
fmt.Println("Starting server...")
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}

	return nil
}
