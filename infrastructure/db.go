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
	DB_URL := os.Getenv("postgres://ffhyfswyqaxkyf:78962915610884f7651e62944feb321e64600786fd02d16eaddfef437a182378@ec2-34-235-198-25.compute-1.amazonaws.com:5432/dcvb1smjrbgapl")

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
