package settings

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/spf13/viper"
)

var Configuration *viper.Viper

func init() {
	environment := os.Getenv("ENVIROMENT")
	if environment == "" {
		os.Setenv("ENVIROMENT", "development")
		fmt.Println(color.Ize(color.Yellow, "environment variable is not set, setting environment to: development"))
	}

	Configuration = viper.New()
	Configuration.SetConfigType("yaml")
	Configuration.SetConfigName(os.Getenv("ENVIROMENT"))
	Configuration.AddConfigPath("../settings")
	Configuration.AddConfigPath("settings/")
	err := Configuration.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
