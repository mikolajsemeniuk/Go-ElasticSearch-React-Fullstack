package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/extensions"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/inputs"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/services"
)

var AccountController IAccountController = &accountController{}

type IAccountController interface {
	FindAccounts(context *gin.Context)
	AddAccount(context *gin.Context)
	FindAccount(context *gin.Context)
	RemoveAccount(context *gin.Context)
	UpdateAccount(context *gin.Context)
}

type accountController struct{}

func (*accountController) FindAccounts(context *gin.Context) {
	payloads, err := services.AccountService.FindAccounts()

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
	context.Next()
}

func (*accountController) AddAccount(context *gin.Context) {
	input := context.MustGet("input").(*inputs.Account)

	payloads, err := services.AccountService.AddAccount(*input)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}

func (*accountController) FindAccount(context *gin.Context) {
	id := context.MustGet("id").(uuid.UUID)

	payloads, err := services.AccountService.FindAccount(id)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}

func (*accountController) RemoveAccount(context *gin.Context) {
	id := context.MustGet("id").(uuid.UUID)

	payloads, err := services.AccountService.RemoveAccount(id)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}

func (*accountController) UpdateAccount(context *gin.Context) {
	id := context.MustGet("id").(uuid.UUID)
	body := context.MustGet("input").(*inputs.Account)

	payloads, err := services.AccountService.UpdateAccount(id, *body)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}
