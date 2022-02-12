package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(context *gin.Context) {
	err := context.MustGet("err")
	payloads := context.MustGet("payloads")

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":   nil,
			"errors": []string{err.(error).Error()},
		})
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data":   payloads,
		"errors": []string{},
	})
}
