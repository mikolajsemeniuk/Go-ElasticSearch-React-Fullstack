package settings

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/spf13/viper"
)

var (
	Configuration *viper.Viper
	Enviroment    string = "development"
)

func init() {
	environment := os.Getenv("ENVIROMENT")

	if environment == "" {
		fmt.Println(color.Ize(color.Yellow, "ENVIROMENT variable is not set, setting environment to: development"))
	}

	if environment != "" {
		Enviroment = environment
	}

	Configuration = viper.New()
	Configuration.SetConfigType("yaml")
	Configuration.SetConfigName(Enviroment)
	Configuration.AddConfigPath("../settings")
	Configuration.AddConfigPath("settings/")
	err := Configuration.ReadInConfig()

	if err != nil {
		panic(color.Ize(color.Red, err.Error()))
	}
}
