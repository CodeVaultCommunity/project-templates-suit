// Package hellosay handler
package hellosay

import (
	"database/sql"
	"errors"
	"fmt"
	errorsapp "mod_name/error"
	"mod_name/repository"
	"mod_name/utils/request"
	"net/http"
	"path"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

func service(c *gin.Context, id int64) *Sample {
	sample := &Sample{}

	_, filename, _, _ := runtime.Caller(0)

	row := repository.FastQueryRow(c, path.Join(path.Dir(filename), "./repo/select_sample.sql"), id)

	if err := row.Scan(&sample.ID, &sample.Name, &sample.Grettings); err != nil {
		panic(err)
	}

	return sample
}

func toPresentation(obj *Sample) *ResponseBody {
	r := &ResponseBody{}

	builder := strings.Builder{}
	builder.WriteString(obj.Grettings)
	builder.WriteString(" ")
	builder.WriteString(obj.Name)
	builder.WriteString("! Your id is ")
	builder.WriteString(fmt.Sprintf("%d", obj.ID))

	r.Message = builder.String()

	return r
}

func handlePanic() {
	if p := recover(); p != nil {
		err, ok := p.(error)
		if ok && errors.Is(err, sql.ErrNoRows) {
			panic(errorsapp.ErrNotFound)
		}
	}
}

// SayHello 	 godoc
// @Summary      Simple test swagger
// @Description  Expect a GET method and returns a simple message
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Param        id   path     int  true  "Wanted ID"  example(20)
// @Router       /helloworld/sayhello/{id} [get]
func SayHello(c *gin.Context) {
	defer handlePanic()

	id := request.GetIdFromUrl(c)
	sample := service(c, id)
	response := toPresentation(sample)

	c.JSON(http.StatusOK, response)
}
