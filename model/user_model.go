package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username    string    `gorm:"unique;not null" json:"username"`
	Password    string    `gorm:"not null"`
	FullName    string    `gorm:"not null" json:"full_name"`
	Email       string    `gorm:"unique;not null" json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `gorm:"type:enum('admin','staff','nasabah');not null" json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}
