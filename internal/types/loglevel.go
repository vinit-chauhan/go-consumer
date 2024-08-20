package types

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel int

const (
	QUITE LogLevel = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

func (l LogLevel) String() string {
	return [...]string{
		"quite",
		"error",
		"warn",
		"info",
		"debug",
	}[l]
}

func (l LogLevel) EnumIndex() int {
	return int(l)
}

func (l LogLevel) ZapLevel() zapcore.Level {
	switch l {
	case QUITE:
		return zap.PanicLevel
	case ERROR:
		return zap.ErrorLevel
	case WARNING:
		return zap.WarnLevel
	case INFO:
		return zap.InfoLevel
	case DEBUG:
		return zap.DebugLevel
	}

	return zap.InfoLevel
}
