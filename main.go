package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	application "github.com/hanif-adedotun/ecommerce-golang/application"
)

func main() {
	// Connect to DB first
	// Connect to routers
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}

}
