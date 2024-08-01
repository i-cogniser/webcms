package config

import (
	"go.uber.org/zap"
	"time"
	"webcms/repositories"
)

func StartTokenCleanupProcess(tokenRepository repositories.TokenRepository, sugar *zap.SugaredLogger) {
	for {
		sugar.Infof("Starting token cleanup process")
		err := tokenRepository.DeleteExpiredTokens()
		if err != nil {
			sugar.Errorf("Failed to delete expired tokens: %v", err)
		} else {
			sugar.Infof("Expired tokens deleted successfully")
		}
		time.Sleep(1 * time.Hour)
	}
}
