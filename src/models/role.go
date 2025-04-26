package models

type Role struct {
	RoleID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
