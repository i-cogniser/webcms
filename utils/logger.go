package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	// Настройка конфигурации логгера
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Создание нового логгера
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic("failed to initialize zap logger")
	}

	// Замена глобального стандартного логгера на Zap
	zap.ReplaceGlobals(Logger)
}
