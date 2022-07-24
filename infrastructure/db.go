package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database  structure
type Database struct {
	DB *gorm.DB
}

// Database initialized and returns mysql  db

func NoteAddDatabase() Database {
	DB_URL := os.Getenv("DB_DEV_URL")

	URL := DB_URL

	fmt.Println(URL)
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}
}
