package models

import (
	"time"

	"github.com/google/uuid"
)

type Debt struct {
	DebtID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"debt_id"`
	Value        float64   `gorm:"not null" json:"value" binding:"required"`
	MaturityDate time.Time `gorm:"not null" json:"maturity_date" binding:"required"`
	UserID       uint      `gorm:"not null" json:"user_id"`
}
