// Package hellowolrdrouter register the routes of hellowolrd module
package hellowolrdrouter

import (
	hellosay "mod_name/module/helloworld/hello/say"

	"github.com/gin-gonic/gin"
)

// Register register all helloworld module endpoints
func Register(router *gin.RouterGroup) {
	helloworld := router.Group("/helloworld")

	helloworld.GET("/sayhello/:id", hellosay.SayHello)
}
