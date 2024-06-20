package main

import (
	"context"
	"fmt"

	"github.com/ngthecoder/orders_api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("failed to start application: %v\n", err)
	}
}
