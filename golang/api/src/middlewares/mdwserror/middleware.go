// Package mdwserror contains middleware functions for handling errors during api execution
package mdwserror

import "github.com/gin-gonic/gin"

// ErrorHandlerMiddlewareGetter returns the `errorHandlerMiddleware`
// It is used to wrapper the two use mode:
//
// - Production for ginmode = Release
//
// - Debug for other gin modes
func ErrorHandlerMiddlewareGetter() func(*gin.Context) {
	if gin.Mode() == gin.ReleaseMode {
		return errorHandlerMiddleware
	}
	return errorHandlerMiddlewareDebug
}
