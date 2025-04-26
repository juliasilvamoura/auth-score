package database

import (
	"log"
	"time"

	"github.com/juliasilvamoura/auth-score/src/models"
	"golang.org/x/crypto/bcrypt"
)

func CleanDB() {
	DB.Exec("TRUNCATE TABLE users CASCADE")
	DB.Exec("TRUNCATE TABLE debts CASCADE")
	DB.Exec("TRUNCATE TABLE roles CASCADE")
}

func SeedDB() {
	log.Println("Iniciando seed do banco de dados...")

	roles := []models.Role{
		{RoleID: 1, Name: "admin"},
		{RoleID: 2, Name: "user"},
	}

	for _, role := range roles {
		if err := DB.Create(&role).Error; err != nil {
			log.Printf("Erro ao criar role %s: %v\n", role.Name, err)
		}
	}

	users := []models.User{
		{
			CPF:       "12345678901",
			Name:      "João Silva",
			BirthDate: time.Date(1990, time.March, 10, 0, 0, 0, 0, time.UTC),
			Email:     "joao@example.com",
			Password:  "senha123",
			RoleID:    2,
		},
		{
			CPF:       "98765432100",
			Name:      "Maria Oliveira",
			BirthDate: time.Date(1985, time.July, 25, 0, 0, 0, 0, time.UTC),
			Email:     "maria@br.experian.com",
			Password:  "admin123",
			RoleID:    1,
		},
	}

	// Array para armazenar os IDs dos usuários criados
	var createdUsers []models.User

	for _, user := range users {
		// Hash da senha antes de salvar
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Erro ao gerar hash da senha para usuário %s: %v\n", user.CPF, err)
			continue
		}
		user.Password = string(hashedPassword)

		if err := DB.Create(&user).Error; err != nil {
			log.Printf("Erro ao criar user %s: %v\n", user.CPF, err)
			continue
		}

		// Adiciona o usuário criado ao array
		createdUsers = append(createdUsers, user)
	}

	// Verifica se os usuários foram criados
	if len(createdUsers) != 2 {
		log.Println("Erro: Nem todos os usuários foram criados corretamente")
		return
	}

	debts := []models.Debt{
		{
			Value:        150.75,
			MaturityDate: time.Now().AddDate(0, 3, 0),
			UserID:       createdUsers[0].UserID,
		},
		{
			Value:        250.00,
			MaturityDate: time.Now().AddDate(0, 6, 0),
			UserID:       createdUsers[1].UserID,
		},
	}

	for _, debt := range debts {
		if err := DB.Create(&debt).Error; err != nil {
			log.Printf("Erro ao criar dívida: %v\n", err)
		}
	}

	log.Println("Seed do banco de dados concluído!")
}
