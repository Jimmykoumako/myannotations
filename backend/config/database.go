// config/config.go
package config

import (
	"fmt"
	"log"
	"os"
	"mas/migrations"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

/// InitDB initializes the database connection and performs migrations
func InitDB() {
	postgresURL := os.Getenv("POSTGRES_URL")

	// If running in a Docker container, use the provided environment variables
	if inDocker := os.Getenv("IN_DOCKER"); inDocker == "true" {
		postgresURL = "postgres://postgres:postgres@database:5432/madb"
	}

	// Connect to the database
	conn, err := gorm.Open("postgres", postgresURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set the global DB variable
	DB = conn

	// Perform database migrations
	if err := migrations.CreateUserTable(DB); err != nil {
		log.Fatal("Failed to migrate User table:", err)
	}

	if err := migrations.CreateBookTable(DB); err != nil {
		log.Fatal("Failed to migrate Book table:", err)
	}

	if err := migrations.CreateTextTable(DB); err != nil {
		log.Fatal("Failed to migrate Text table:", err)
	}

	if err := migrations.CreateAnnotationTable(DB); err != nil {
		log.Fatal("Failed to migrate Annotation table:", err)
	}

	if err := migrations.CreateConnectionTable(DB); err != nil {
		log.Fatal("Failed to migrate Connection table:", err)
	}

	if err := migrations.CreateConnectionFeedbackTable(DB); err != nil {
		log.Fatal("Failed to migrate ConnectionFeedback table:", err)
	}

	fmt.Println("Database initialization and migrations completed")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Closed the database connection")
	}
}