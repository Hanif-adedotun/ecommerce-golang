package main

import (
	"context"
	application "github.com/hanif-adedotun/ecommerce-golang/application"
	"fmt"
)

func main() {
	// Connect to DB first
	// Connect to routers
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}

}
