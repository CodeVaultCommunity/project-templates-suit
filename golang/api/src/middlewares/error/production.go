// Package middleware_errorhandler contains middleware functions for handling errors during api execution
package middleware_errorhandler

import (
	"log"
	errorsapp "mod_name/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleAppError(c *gin.Context, e *errorsapp.AppError) {
	log.Printf("[WARN] %s: %v", e.Code, e.Err)
	c.JSON(e.Status, gin.H{"error": e.Message, "code": e.Code})
}

func handleInternalServerError(c *gin.Context, e error) {
	log.Printf("[PANIC] %v", e)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}

func handleUnforeseenError(c *gin.Context, r any) {
	log.Printf("[PANIC] %v", r)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Unspected internal server error"})
}

func handleError(c *gin.Context) {
	if r := recover(); r != nil {
		switch e := r.(type) {
		case *errorsapp.AppError:
			handleAppError(c, e)
		case error:
			handleInternalServerError(c, e)
		default:
			handleUnforeseenError(c, r)
		}
		c.Abort()
	}
}

func errorHandlerMiddleware(c *gin.Context) {
	defer handleError(c)
	c.Next()
}
