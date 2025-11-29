// Package api describes the purpose of the module.
package api

import (
	"log"
	middleware_errorhandler "mod_name/middlewares/error"
	"os"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine = nil
var api *gin.RouterGroup = nil

func startEngine() *gin.Engine {
	if engine != nil {
		return engine
	}

	engine = gin.Default()
	engine.Use(middleware_errorhandler.ErrorHandlerMiddlewareGetter())

	return engine
}

func StartAPI() *gin.RouterGroup {
	if api != nil {
		return api
	}

	startEngine()

	api = engine.Group(os.Getenv("SUBPATH"))
	return api
}

func RunServer() {
	os.Getenv("")
	if err := engine.Run(os.Getenv("API_PORT")); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
