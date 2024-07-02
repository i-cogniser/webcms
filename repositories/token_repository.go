package repositories

import (
	"time"
	"webcms/models"

	"github.com/jinzhu/gorm"
)

type TokenRepository interface {
	SaveToken(token models.Token) error
	SaveTokenWithTx(token models.Token, tx *gorm.DB) error
	GetToken(tokenString string) (*models.Token, error)
	DeleteToken(tokenString string) error
	DeleteExpiredTokens() error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db}
}

func (r *tokenRepository) SaveToken(token models.Token) error {
	return r.db.Create(&token).Error
}

func (r *tokenRepository) SaveTokenWithTx(token models.Token, tx *gorm.DB) error {
	return tx.Create(&token).Error
}

func (r *tokenRepository) GetToken(tokenString string) (*models.Token, error) {
	var token models.Token
	err := r.db.Where("token = ?", tokenString).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) DeleteToken(tokenString string) error {
	return r.db.Where("token = ?", tokenString).Delete(&models.Token{}).Error
}

func (r *tokenRepository) DeleteExpiredTokens() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&models.Token{}).Error
}
