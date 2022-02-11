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
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		context.Set("input", input)
	}
}
