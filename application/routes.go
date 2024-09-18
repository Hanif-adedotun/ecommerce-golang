package application

import (
	"net/http"

	"github.com/hanif-adedotun/ecommerce-golang/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoute)

	return router
}

func loadOrderRoute(router chi.Router) {
	orderHandler := &handler.Order{}

	// All routes for the order handler 
	// All path is orders/
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetbyID)
	router.Put("/{id}", orderHandler.UpdatebyID)
	router.Delete("/{id}", orderHandler.DeletebyID)
}
