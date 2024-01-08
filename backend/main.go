package main

import (
	"mas/config"
	"mas/migrations"
	"mas/routers"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Defer closing the database connection
	defer config.DB.Close()

	// Apply migrations
	migrations.RunMigrations()

	// Initialize routes and start the server
	router := routers.InitializeRoutes()
	router.Run(":8080")
}
