package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/controllers"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/inputs"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/middlewares"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/settings"
)

var router = gin.Default()

func Listen() {
	v1 := router.Group(settings.Configuration.GetString("server.basepath"))
	{
		accounts := v1.Group("accounts")
		{
			accounts.GET("", controllers.AccountController.FindAccounts, middlewares.Response)
			accounts.POST("", middlewares.Body(inputs.Account{}), controllers.AccountController.AddAccount, middlewares.Response)
			accounts.Use(middlewares.Route("id", uuid.UUID{}))
			accounts.GET(":id", controllers.AccountController.FindAccount, middlewares.Response)
			accounts.DELETE(":id", controllers.AccountController.RemoveAccount, middlewares.Response)
			accounts.PATCH(":id", middlewares.Body(inputs.Account{}), controllers.AccountController.UpdateAccount, middlewares.Response)
		}
	}
	router.Use(middlewares.Response)
	port := fmt.Sprintf(":%s", settings.Configuration.GetString("server.port"))
	router.Run(port)
}
