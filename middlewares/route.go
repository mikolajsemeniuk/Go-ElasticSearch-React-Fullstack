package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Route(name string, kind interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {

		var (
			err   error
			param interface{}
		)

		switch kind.(type) {
		case uuid.UUID:
			param, err = uuid.Parse(context.Param(name))
		case int:
			param, err = strconv.Atoi(context.Param(name))
		}

		fmt.Println(param)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"data":    nil,
				"errors":  []string{err.Error()},
				"message": fmt.Sprintf("%s not valid", name),
			})
			return
		}

		context.Set(name, param)
	}
}
