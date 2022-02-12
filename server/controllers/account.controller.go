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

// @Summary get all accounts
// @Schemes
// @Description get all accounts
// @Tags accounts
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /accounts [get]
func (*accountController) FindAccounts(context *gin.Context) {
	payloads, err := services.AccountService.FindAccounts()

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
	context.Next()
}

// @Summary add account
// @Schemes
// @Description add account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body inputs.Account true "account to create"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts [post]
func (*accountController) AddAccount(context *gin.Context) {
	input := context.MustGet("input").(*inputs.Account)

	payloads, err := services.AccountService.AddAccount(*input)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}

// @Summary get account by id
// @Schemes
// @Description get account by id
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts/{accountId} [get]
func (*accountController) FindAccount(context *gin.Context) {
	id := context.MustGet("id").(uuid.UUID)

	payloads, err := services.AccountService.FindAccount(id)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}

// @Summary remove account
// @Schemes
// @Description remove account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts/{accountId} [delete]
func (*accountController) RemoveAccount(context *gin.Context) {
	id := context.MustGet("id").(uuid.UUID)

	payloads, err := services.AccountService.RemoveAccount(id)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}

// @Summary update account
// @Schemes
// @Description update account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param account body inputs.Account true "account to update"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts/{id} [patch]
func (*accountController) UpdateAccount(context *gin.Context) {
	id := context.MustGet("id").(uuid.UUID)
	body := context.MustGet("input").(*inputs.Account)

	payloads, err := services.AccountService.UpdateAccount(id, *body)

	context.Set("payloads", payloads)
	context.Set("err", err)

	extensions.Info("done")
}
