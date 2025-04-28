package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/juliasilvamoura/auth-score/src/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	fmt.Println("Conectando ao banco de dados...")
	database.ConnectDB()

	fmt.Println("Limpando dados existentes...")
	database.CleanDB()

	fmt.Println("Populando banco de dados com dados de teste...")
	database.SeedDB()

	fmt.Println("Processo concluído!")
}
