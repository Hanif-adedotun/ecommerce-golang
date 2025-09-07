package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hanif-adedotun/ecommerce-golang/db"
	"github.com/hanif-adedotun/ecommerce-golang/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func TestLoadRoutes(t *testing.T) {
	// Arrange
	a := &App{}
	a.db, _ = sql.Open("pgx", "")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	// Act
	a.loadRoutes()
	a.router.ServeHTTP(w, r)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestLoadOrderRoute(t *testing.T) {
	// Arrange
	a := &App{}
	a.db, _ = sql.Open("pgx", "")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/orders", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	// Act
	router := chi.NewRouter()
	a.loadOrderRoute(router)
	router.ServeHTTP(w, r)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
