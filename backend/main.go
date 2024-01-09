package main

import (
	"mas/config"
	"mas/routers"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Defer closing the database connection
	defer config.DB.Close()

	

	// Initialize routes and start the server
	routers.Router.Run(":8080")
}
