package main

import (
	"go-rest-backend/db"
	"go-rest-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize database
	db.InitDB()

	// Create server
	server := gin.Default()

	// Register routes
	routes.RegisterRoutes(server)

	// Start server
	server.Run(":8080")
}