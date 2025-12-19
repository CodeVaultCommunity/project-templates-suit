// Package repository - includes all needed to initialize the database
package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver

	// _ "github.com/lib/pq"                 // PostgreSQL driver
	"fmt"
	"log"
	"mod_name/config"
)

// SQLDB is the database object
var SQLDB *sql.DB

// DatabaseInit initialize the database
func DatabaseInit() error {
	db, err := sql.Open(config.Engine, config.DatabaseDNS)

	if err != nil {
		return fmt.Errorf("can't connect with database: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return fmt.Errorf("error try ping database: %w", err)
	}

	SQLDB = db
	log.Println("Database connection success")
	return nil
}

// DatabaseClose closes the database
func DatabaseClose() {
	if err := SQLDB.Close(); err != nil {
		log.Printf("error closing database connection: %v", err)
	}
}
