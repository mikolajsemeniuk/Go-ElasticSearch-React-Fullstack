package extensions

import (
	"github.com/TwiN/go-color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	encoderConfiguration := zap.NewDevelopmentEncoderConfig()
	encoderConfiguration.TimeKey = "timestamp"
	encoderConfiguration.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfiguration.EncodeLevel = zapcore.CapitalColorLevelEncoder

	configuration := zap.NewDevelopmentConfig()
	configuration.EncoderConfig = encoderConfiguration

	var err error
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
