// Package config provides utility functions for loading and configuring
// environment variables, API settings, and Swagger documentation for the project instance.
package config

import (
	"fmt"
	"os"
)

// setenv checks whether an environment variable is already defined.
// If not, it sets it to the provided default value.
//
// Returns the final value and any error encountered when setting the variable.
func setenv(key string, defaultValue string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		err := os.Setenv(key, defaultValue)
		return defaultValue, err
	}

	return val, nil
}

func checkNoOptionalKey(key string, canBeEmpty bool) error {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fmt.Errorf("the key %s is not optional", key)
	}

	if !canBeEmpty && len(val) == 0 {
		return fmt.Errorf("the key %s can't be empty", key)
	}

	return nil
}
