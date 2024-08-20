package utils

import (
	"github.com/vinit-chauhan/go-consumer/internal/types"
	"go.uber.org/zap"
)

func Logger(conf types.Config) (*zap.Logger, error) {
	level := conf.LogLevel.ZapLevel()

	logger := zap.NewProductionConfig()
	logger.OutputPaths = []string{"service.log"}
	logger.Encoding = "console"
	logger.Level.SetLevel(level)

	return logger.Build()
}
