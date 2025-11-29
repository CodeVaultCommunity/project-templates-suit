package main

import (
	"log"
	"mod_name/api"
	"mod_name/config"
	"mod_name/repository"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := repository.DatabaseInit(); err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	// Close the database connection when main finishes
	defer repository.DatabaseClose()

	api.StartAPI()

	api.RunServer()
}
