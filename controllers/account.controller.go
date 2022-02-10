package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	accounts, err := services.AccountService.FindAccounts()
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, gin.H{
			"data":    nil,
			"errors":  []string{"database error"},
			"message": "All accounts were fetched",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data":    accounts,
		"errors":  []string{},
		"message": "All accounts were fetched",
	})
}

func (*accountController) FindAccountById(context *gin.Context) {

}

func (*accountController) AddAccount(context *gin.Context) {

}

func (*accountController) UpdateAccount(context *gin.Context) {

}

func (*accountController) RemoveAccount(context *gin.Context) {

}
