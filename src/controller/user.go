package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juliasilvamoura/auth-score/src/database"
	"github.com/juliasilvamoura/auth-score/src/models"
)

func PostUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos ou incompletos"})
		return
	}

	atIndex := strings.Index(user.Email, "@")
	if atIndex == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email inválido"})
		return
	}
	domain := user.Email[atIndex:]
	if domain == "@br.experian.com" {
		user.RoleID = 1
	} else {
		user.RoleID = 2
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := database.DB.First(&existingUser, user.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	oldRoleID := existingUser.RoleID

	if err := database.DB.Model(&existingUser).Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Se o role foi alterado, invalida os tokens existentes
	if oldRoleID != user.RoleID {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			// Adiciona o token atual à blacklist com o tempo restante de expiração
			database.AddToBlacklist(tokenString, 0) // 0 para usar o tempo padrão de expiração
		}
	}

	c.JSON(http.StatusOK, existingUser)
}
