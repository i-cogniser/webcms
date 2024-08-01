package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadEnv(sugar *zap.SugaredLogger) {
	if err := godotenv.Load(); err != nil {
		sugar.Fatalf("Error loading .env file: %v", err)
	}
	sugar.Infof(".env file loaded successfully")
}
