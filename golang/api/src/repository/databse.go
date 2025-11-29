// Package repository - includes all needed to initialize the database
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"mod_name/config"
)

// DB is the database
var DB *sql.DB

func DatabaseInit() error {
	db, err := sql.Open("postgres", config.DATABASE_URL)

	if err != nil {
		return fmt.Errorf("can't connect with database: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return fmt.Errorf("error try ping database: %w", err)
	}

	DB = db
	log.Println("Database connection success")
	return nil
}

func DatabaseClose() {
	if err := DB.Close(); err != nil {
		log.Printf("error closing database connection: %v", err)
	}
}
