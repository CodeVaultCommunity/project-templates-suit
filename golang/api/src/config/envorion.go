// Package config provides utility functions for loading and configuring
// environment variables, API settings, and Swagger documentation for the project instance.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func checkNoOptinalAPIKeys() error {
	var err error

	err = checkNoOpitionalKey("DATABASE_URL", false)
	if err != nil {
		return err
	}

	return nil
}

// checkAndSetOpitionalAPIKeys defines and validates the basic runtime environment variables
// required for the API to operate properly.
//
// It ensures:
//   - API_PORT is set (defaults to ":8080")
//   - GIN_MODE is set and valid (defaults to gin.DebugMode)
//
// If GIN_MODE is set to an invalid value, it automatically falls back to gin.DebugMode.
func checkAndSetOpitionalAPIKeys() error {
	setValue, err := setenv("API_PORT", ":8080")
	if err != nil {
		return err
	}

	if setValue[0] != ':' {
		os.Setenv("API_PORT", fmt.Sprintf(":%s", setValue))
	}

	setValue, err = setenv("GIN_MODE", gin.DebugMode)
	ginModeIsValid := setValue == gin.DebugMode || setValue == gin.TestMode || setValue == gin.ReleaseMode

	if !ginModeIsValid {
		// If GIN_MODE was defined in .env but contains an invalid value,
		// revert to the default mode (Debug).
		log.Printf("The value %s is not valid for GIN_MODE", setValue)
		err = os.Setenv("GIN_MODE", gin.DebugMode)
	}

	if err != nil {
		return err
	}

	_, err = setenv("SUBPATH", "/")
	if err != nil {
		return err
	}

	return nil
}

// Load initializes all configuration settings required for the project runtime.
//
// It performs the following steps:
//  1. Loads environment variables from a .env file.
//  2. Validates and sets API configuration (port and mode).
//  3. Applies the configured Gin mode.
//
// Returns an error if any step fails.
func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: failed to load .env file:", err)
		return err
	}

	err = checkNoOptinalAPIKeys()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = checkAndSetOpitionalAPIKeys()
	if err != nil {
		log.Println("Warning: failed to set API configuration:", err)
		return err
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	err = loadConstraints()

	return err
}
