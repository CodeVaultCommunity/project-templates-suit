// Package main simple the main function
package main

import (
	"fmt"
	"log"
	"mod_name/api"
	"mod_name/config"
	_ "mod_name/docs" // import swagger docs
	hellowolrdrouter "mod_name/module/helloworld/.router"
	"mod_name/repository"
	"os"

	"github.com/gin-gonic/gin"
)

func runServer(engine *gin.Engine) error {
	if err := engine.Run(os.Getenv("API_PORT")); err != nil {
		return fmt.Errorf("error starting server: %v", err)
	}

	return nil
}

// @title API Title
// @version 1.0
// @description ...
// @host localhost:8080
// @BasePath /
func main() {
	if err := config.Load(true); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := repository.DatabaseInit(); err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	// Close the database connection when main finishes
	defer repository.DatabaseClose()

	engine := api.StartEngine()
	router := api.StartAPI()

	if router == nil {
		log.Println("engine is not started yet, runs StartEngine function")
		return
	}

	hellowolrdrouter.Register(router)

	if err := runServer(engine); err != nil {
		log.Printf("Fatal error: %v", err)
		return
	}
}
