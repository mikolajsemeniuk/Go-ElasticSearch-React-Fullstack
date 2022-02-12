package extensions

import (
	"os"

	"github.com/TwiN/go-color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	err    error
)

func init() {
	encoderConfiguration := zap.NewDevelopmentEncoderConfig()
	configuration := zap.NewDevelopmentConfig()

	environment := os.Getenv("ENVIROMENT")
	if environment != "development" {
		encoderConfiguration = zap.NewProductionEncoderConfig()
		configuration = zap.NewProductionConfig()
	}

	encoderConfiguration.TimeKey = "timestamp"
	encoderConfiguration.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfiguration.EncodeLevel = zapcore.CapitalColorLevelEncoder

	configuration.EncoderConfig = encoderConfiguration

	logger, err = configuration.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	logger.Info(color.Ize(color.Green, message), fields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(color.Ize(color.Yellow, message), fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(color.Ize(color.Red, message), fields...)
}
