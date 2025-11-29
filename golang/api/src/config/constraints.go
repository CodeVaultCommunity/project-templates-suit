// Package config provides utility functions for loading and configuring
// environment variables, API settings, and Swagger documentation for the project instance.
package config

import (
	"errors"
	"os"
)

var (
	DATABASE_URL string
)

func loadConstraints() error {
	var ok bool

	DATABASE_URL, ok = os.LookupEnv("DATABASE_URL")

	if !ok {
		return errors.New("cant load DATABASE_URL from envoriment")
	}

	return nil
}
