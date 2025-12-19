// Package mdwserror contains middleware functions for handling errors during api execution
package mdwserror

import (
	"errors"
	"fmt"
	"log"
	errorsapp "mod_name/error"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func contextDebugger(c *gin.Context, err error) map[string]string {
	msg := "<nil>"
	if err != nil {
		msg = err.Error()
	}

	return map[string]string{
		"status":     "error",
		"message":    msg,
		"stacktrace": string(debug.Stack()),
		"path":       c.FullPath(),
		"method":     c.Request.Method,
		"client_ip":  c.ClientIP(),
		"ATTENTION": fmt.Sprintf(
			"You are seeing this because ginmode=%s Switch to ginmode=%s when moving to production",
			gin.Mode(),
			gin.ReleaseMode,
		),
	}
}

func handleAppErrorDebug(c *gin.Context, e *errorsapp.AppError) {
	log.Printf("[WARN] %s: %v", e.Code, e.Err)
	c.JSON(e.Status, gin.H{"error": e.Message, "code": e.Code, "debug": contextDebugger(c, e.Err)})
}

func handleInternalServerErrorDebug(c *gin.Context, e error) {
	log.Printf("[PANIC] %v", e)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "debug": contextDebugger(c, e)})
}

func handleUnforeseenErrorDebug(c *gin.Context, r any) {
	log.Printf("[PANIC] %v", r)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Unspected internal server error", "debug": contextDebugger(c, errors.New("no error provided")), "r": r})
}

func handleErrorDebug(c *gin.Context) {
	if r := recover(); r != nil {
		switch e := r.(type) {
		case *errorsapp.AppError:
			handleAppErrorDebug(c, e)
		case error:
			handleInternalServerErrorDebug(c, e)
		default:
			handleUnforeseenErrorDebug(c, r)
		}
		c.Abort()
	}
}

func errorHandlerMiddlewareDebug(c *gin.Context) {
	defer handleErrorDebug(c)
	c.Next()
}
