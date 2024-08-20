package utils

import (
	"github.com/vinit-chauhan/go-consumer/internal/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func Init(conf types.Config) error {
	level := conf.LogLevel.ZapLevel()

	logger := zap.NewProductionConfig()

	logger.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger.OutputPaths = []string{"service.log"}
	logger.Encoding = "console"
	logger.Level.SetLevel(level)

	l, err := logger.Build()
	if err != nil {
		return err
	}
	defer l.Sync()

	Logger = l.Sugar()

	return nil
}
