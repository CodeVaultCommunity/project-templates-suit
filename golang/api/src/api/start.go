// Package api contains some utils functions to start the api
package api

import (
	"mod_name/config"
	middleware_errorhandler "mod_name/middlewares/mdwserror"
	"os"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine
var api *gin.RouterGroup

// StartEngine initialize a default gin engine with the error handler middleware and swagger registered
func StartEngine() *gin.Engine {
	if engine != nil {
		return engine
	}

	engine = gin.Default()
	engine.Use(middleware_errorhandler.ErrorHandlerMiddlewareGetter())
	config.RegisterSwagger(engine)

	return engine
}

// StartAPI starts the default gin engine, configure a group with the SUBPATH variable and returns this group.
// Capture this group and use it to register the routes and methods.
func StartAPI() *gin.RouterGroup {
	if engine == nil {
		return nil
	}

	if api != nil {
		return api
	}

	api = engine.Group(os.Getenv("SUBPATH"))

	return api
}
