package database

import (
	"fmt"
	"log"
	"strconv"

	"example.org/config"
	"example.org/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// Connect to the database
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Error parsing port")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect to the database!")
	}
	defer func() {
		dbInstance, _ := DB.DB()
		_ = dbInstance.Close()
	}()

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")
}
