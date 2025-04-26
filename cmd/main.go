package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rraj7/demo-fintech-app/internal/adapters/http"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	http.RegisterRoutes(router)
	router.Run(":8080") // Start the server on port 8080
}

// This is a simple main function that initializes a Gin router and registers routes for the application.
