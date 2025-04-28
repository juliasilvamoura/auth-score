package database

import (
	"log"

	"github.com/juliasilvamoura/auth-score/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=auth_score port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Enable UUID extension
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal("Error enabling UUID extension:", err)
	}

	DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Debt{},
	)

	DB.Exec("SET CONSTRAINTS ALL IMMEDIATE")
}
