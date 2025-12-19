// Package request provides some utilit functions
package request

import (
	errorsapp "mod_name/error"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetNotOptionalParam gets the passed key from url query and checks if it is empty
func GetNotOptionalParam(c *gin.Context, key string) string {
	value := c.Param(key)
	if len(value) == 0 {
		panic(errorsapp.ErrNotOptionalURLParam(key))
	}

	return value
}

func GetIdFromUrl(c *gin.Context) int64 {
	id := GetNotOptionalParam(c, "id")

	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errApp := errorsapp.ErrURLParamInvalidType("id", "integer")
		errApp.Err = err
		panic(errApp)
	}

	return idNumber
}
