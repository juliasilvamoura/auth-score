package database

import (
	"log"
	"os"

	"github.com/juliasilvamoura/auth-score/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := os.Getenv("URL_DATABASE")
	if dsn == "" {
		log.Fatal("URL_DATABASE environment variable is not set")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Auto Migrate in correct order
	DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Debt{},
	)

	// Enable foreign keys after migration
	DB.Exec("SET CONSTRAINTS ALL IMMEDIATE")
}
