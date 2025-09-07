package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hanif-adedotun/ecommerce-golang/db"
	"github.com/hanif-adedotun/ecommerce-golang/handler"
)

func TestAppStart(t *testing.T) {
	// Arrange
	ctx := context.Background()
	app := &App{}
	app.db, _ = sql.Open("pgx", "host=localhost user=myuser password=mypassword port=5432 database=mydb")
	app.loadRoutes()

	// Act
	if err := app.Start(ctx); err != nil {
		t.Errorf("error starting app: %v", err)
	}
}

func TestLoadRoutes(t *testing.T) {
	// Arrange
	app := &App{}
	app.loadRoutes()

	// Assert
	if app.router == nil {
		t.Errorf("router is nil after loading routes")
	}
}
