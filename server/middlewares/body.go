package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/inputs"
)

func Body(input interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {

		switch input.(type) {
		case inputs.Account:
			input = &inputs.Account{}
		}

		if err := context.BindJSON(&input); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"data":    nil,
				"errors":  []string{err.Error()},
				"message": "body is not a valid JSON",
			})
			return
		}

		context.Set("input", input)
	}
}
