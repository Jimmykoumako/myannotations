// main.go

package main

import (
	// "github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"mas/config"
	"mas/controllers"
	"mas/models"
	"mas/routers"
)

func main() {
	// Initialize the Gin router
	router := routers.SetupRouter()

	// Initialize the database connection
	db, err := config.GetDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	defer db.Close()

	// AutoMigrate your models
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Annotation{}, &models.Connection{}, &models.Text{})

	// Inject the database instance into the controllers
	controllers.SetDB(db)

	// Run the Gin application
	router.Run(":8080")
}
