package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/juliasilvamoura/auth-score/src/database"
	"github.com/juliasilvamoura/auth-score/src/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	jwt_key, exists := c.Get("JWT_SECRET")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET não encontrado"})
		c.Abort()
		return
	}

	jwtKeyBytes, ok := jwt_key.([]byte)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET inválido"})
		c.Abort()
		return
	}

	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var user models.User
	if err := database.DB.Where("cpf = ?", creds.CPF).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "CPF ou senha incorretos"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "CPF ou senha incorretos"})
		return
	}

	expirationTime := time.Now().Add(20 * time.Minute)

	claims := &models.Claims{
		UserID: user.UserID,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeyBytes)
	if err != nil {
		fmt.Println("Erro ao gerar token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível gerar o token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout efetuado. Apague o token no cliente!"})
}
