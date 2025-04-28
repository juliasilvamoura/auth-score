package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juliasilvamoura/auth-score/src/database"
	"github.com/juliasilvamoura/auth-score/src/middleware"
	"github.com/juliasilvamoura/auth-score/src/routes"
)

func main() {
	log.Println("Iniciando o servidor da API")

	_ = godotenv.Load()

	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET não encontrado nas variáveis de ambiente")
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Conectar ao banco de dados
	database.ConnectDB()

	// Conectar ao Redis
	database.ConnectRedis()

	auth := r.Group("/auth")

	auth.Use(func(c *gin.Context) {
		c.Set("JWT_SECRET", jwtKey)
		c.Next()
	})

	auth.Use(func(c *gin.Context) {
		if c.FullPath() != "/auth/login" {
			middleware.AuthMiddleware()(c)
		}
	})

	routes.HandleRequests(r, auth)

	r.Run(":8080")
}
