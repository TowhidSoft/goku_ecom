package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/middleware"
)

func Serve() {

	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)
	gobalRouter := middleware.CorsWithPreflight(mux)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", gobalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
