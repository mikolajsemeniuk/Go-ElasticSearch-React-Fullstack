package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
			accounts.GET("", controllers.AccountController.FindAccounts)
			accounts.POST("", middlewares.Body(inputs.Account{}), controllers.AccountController.AddAccount)
		}
	}
	port := fmt.Sprintf(":%s", settings.Configuration.GetString("server.port"))
	router.Run(port)
}
