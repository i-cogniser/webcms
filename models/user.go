package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `json:"username" validate:"required,min=3" gorm:"unique;not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      string `json:"role" validate:"required,oneof=admin user"`
	Name      string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
