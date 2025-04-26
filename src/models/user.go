package models

import (
	"time"
)

type User struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	CPF       string    `gorm:"unique;not null" json:"cpf" binding:"required"`
	Name      string    `gorm:"not null" json:"name" binding:"required"`
	BirthDate time.Time `gorm:"not null" json:"birth_date" binding:"required"`
	Email     string    `gorm:"unique;not null" json:"email" binding:"required,email"`
	Password  string    `gorm:"not null" json:"password" binding:"required"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
}
