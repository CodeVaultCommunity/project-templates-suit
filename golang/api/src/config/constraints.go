// Package config provides utility functions for loading and configuring
// environment variables, API settings, and Swagger documentation for the project instance.
package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	// DatabaseDNS Configure te database main URL
	DatabaseDNS string
	// Engine configure the API main engine (mysql, postegre, etc)
	Engine string
)

func buildDatabaseDNS() error {
	var ok bool

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return errors.New("can't find the variable DB_USER")
	}
	pass, ok := os.LookupEnv("DB_USER_PASSWORD")
	if !ok {
		return errors.New("can't find the variable DB_USER_PASSWORD")
	}
	name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return errors.New("can't find the variable DB_NAME")
	}
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return errors.New("can't find the variable DB_HOST")
	}
	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return errors.New("can't find the variable DB_PORT")
	}

	switch Engine {
	case "postgre":
		fallthrough
	case "postgres":
		fallthrough
	case "postgresql":
		DatabaseDNS = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
	case "mysql":
		DatabaseDNS = fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Engine, user, pass, host, port, name)
	default:
		return fmt.Errorf("the engine (%s) has no default dns constructor, set the DATABASE_URL envoriment variable or define a constructor for it at file 'src/config/constraints.go'; you will need import the the engine driver at file src/repository/database.go", Engine)
	}

	return nil
}

func loadConstraints() error {
	var ok bool

	Engine, ok = os.LookupEnv("ENGINE")

	if !ok {
		return errors.New("cant load ENGINE from envoriment")
	}

	DatabaseDNS, ok = os.LookupEnv("DATABASE_URL")

	if !ok || len(DatabaseDNS) <= 0 {
		err := buildDatabaseDNS()
		if err != nil {
			return err
		}
	}

	return nil
}
