package main

import (
	"context"
	application "github.com/hanif-adedotun/ecommerce-golang/application"
	db "github.com/hanif-adedotun/ecommerce-golang/db"
	"fmt"
)

func main() {
	// Connect to DB first
	db.ConnectDB()
	// Connect to routers
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
	
}
