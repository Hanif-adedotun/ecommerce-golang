package application

import (
	"net/http"

	"github.com/hanif-adedotun/ecommerce-golang/db"
	"github.com/hanif-adedotun/ecommerce-golang/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes(){
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.loadOrderRoute)

	a.router = router
}

func (a *App) loadOrderRoute(router chi.Router) {
	orderHandler := &handler.Order{
		Repo : &db.PostgreRepo{
			Client: a.db,
		},
	}

	// All routes for the order handler 
	// All path is orders/
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetbyID)
	router.Put("/{id}", orderHandler.UpdatebyID)
	router.Delete("/{id}", orderHandler.DeletebyID)
}
