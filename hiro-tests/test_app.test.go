package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/hanif-adedotun/ecommerce-golang/db"
)

func TestNewApp(t *testing.T) {
	// Arrange
	dbConn, err := sql.Open("pgx", "host=localhost user=postgres password=postgres port=5432 database=postgres")
	if err != nil {
		t.Fatal(err)
	}
	defer dbConn.Close()

db.ConnectDB = func() (*sql.DB, error) {
		return dbConn, nil
	}

	// Act
	app := New()

	// Assert
	if app.db == nil {
		t.Errorf("expected app.db to be not nil, but got nil")
	}
}

func TestStartServer(t *testing.T) {
	// Arrange
	dbConn, err := sql.Open("pgx", "host=localhost user=postgres password=postgres port=5432 database=postgres")
	if err != nil {
		t.Fatal(err)
	}
	defer dbConn.Close()

db.ConnectDB = func() (*sql.DB, error) {
		return dbConn, nil
	}

	app := New()

	// Act and Assert
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}
