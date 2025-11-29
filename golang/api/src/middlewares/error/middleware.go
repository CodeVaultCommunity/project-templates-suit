// Package middleware_errorhandler contains middleware functions for handling errors during api execution
package middleware_errorhandler

import "github.com/gin-gonic/gin"

func ErrorHandlerMiddlewareGetter() func(*gin.Context) {
	if gin.Mode() == gin.ReleaseMode {
		return errorHandlerMiddleware
	}
	return errorHandlerMiddlewareDebug
}
