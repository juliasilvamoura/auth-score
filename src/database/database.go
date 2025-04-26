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
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=root password=root dbname=auth_score port=5432 sslmode=disable"
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Auto Migrate na ordem correta
	DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Debt{},
	)

	// Habilita as foreign keys após a migração
	DB.Exec("SET CONSTRAINTS ALL IMMEDIATE")
}
