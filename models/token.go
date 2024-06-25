package models

import "time"

type Token struct {
	ID        uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"index"`
	Token     string `gorm:"type:text"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
