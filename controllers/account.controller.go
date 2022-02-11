package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/extensions"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/inputs"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/payloads"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/services"
)

var AccountController IAccountController = &accountController{}

type IAccountController interface {
	FindAccounts(context *gin.Context)
	FindAccountById(context *gin.Context)
	AddAccount(context *gin.Context)
	UpdateAccount(context *gin.Context)
	RemoveAccount(context *gin.Context)
}

type accountController struct{}

func (*accountController) FindAccounts(context *gin.Context) {
	payloads, err := services.AccountService.FindAccounts()
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, gin.H{
			"data":    payloads,
			"errors":  []string{err.Error()},
			"message": "error occured",
		})
		return
	}

	extensions.Info("done")
	context.JSON(http.StatusOK, gin.H{
		"data":    payloads,
		"errors":  []string{},
		"message": "All accounts were fetched",
	})
}

func (*accountController) FindAccountById(context *gin.Context) {

}

func (*accountController) AddAccount(context *gin.Context) {
	input := context.MustGet("input").(*inputs.Account)
	payload, err := services.AccountService.AddAccount(*input)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, gin.H{
			"data":    nil,
			"errors":  []string{err.Error()},
			"message": "All accounts were fetched",
		})
		return
	}

	extensions.Info("done")
	context.JSON(http.StatusOK, gin.H{
		"data":    []*payloads.Account{payload},
		"errors":  []string{},
		"message": "Account successfully added",
	})
}

func (*accountController) UpdateAccount(context *gin.Context) {

}

func (*accountController) RemoveAccount(context *gin.Context) {

}
