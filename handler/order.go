package handler

import (
	"fmt"
	"net/http"
)

// This package handles all the routing for orders
// Creating, listing, updating and deleting orders

type Order struct{}

func(o *Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create an order")
}

func(o *Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("List all orders")
}

func(o *Order) GetbyID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get an order by ID")
}

func(o *Order) UpdatebyID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update an order by ID")
}

func(o *Order) DeletebyID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete an order by ID")
}