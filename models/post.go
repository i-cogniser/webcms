package models

import "time"

type Post struct {
	ID        uint   `gorm:"primary_key"`
	Title     string `gorm:"type:varchar(100);not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
