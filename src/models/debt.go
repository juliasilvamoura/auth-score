package models

import (
	"time"

	"github.com/google/uuid"
)

type Debt struct {
	DebtID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"debt_id"`
	Value        float64   `gorm:"not null" json:"value" binding:"required"`
	MaturityDate time.Time `gorm:"not null" json:"maturity_date" binding:"required"`
	UserID       uint      `gorm:"not null" json:"user_id"`
}

type DebtResponse struct {
	DebtID       uuid.UUID `json:"debt_id"`
	Value        float64   `json:"value"`
	MaturityDate time.Time `json:"maturity_date"`
}
