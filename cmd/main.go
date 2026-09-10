package main

import (
	"CRUD_GO/config"
	"CRUD_GO/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database and perform automigration
	config.ConnectDatabase()

	// Seed the database with initial data
	if err := config.RunMigrations(); err != nil {
		log.Fatal("❌ Error en migraciones:", err)
	}
	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")

}
