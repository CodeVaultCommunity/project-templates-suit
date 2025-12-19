// Package config provides utility functions for loading and configuring
// environment variables, API settings, and Swagger documentation for the project instance.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterSwagger registers Swagger documentation routes into the provided Gin router group.
//
// The Swagger UI is only registered when the Gin mode is not set to "release"
// (i.e., in development or test environments). This prevents exposing documentation in production.
//
// Example:
//
//	router := gin.Default()
//	api := router.Group("/api")
//	config.RegisterSwagger(api)
func RegisterSwagger(engine *gin.Engine) {
	if engine == nil {
		return
	}

	// Only expose Swagger documentation in non-production environments.
	if gin.Mode() != gin.ReleaseMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		message := fmt.Sprintf("Swagger successfully registered at: http://localhost%s/swagger/index.html", os.Getenv("API_PORT"))
		log.Println(message)
	}
}
