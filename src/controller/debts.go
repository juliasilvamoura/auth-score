package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juliasilvamoura/auth-score/src/database"
	"github.com/juliasilvamoura/auth-score/src/models"
)

func GetAllDebts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userIDInt, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao extrair o user_id"})
		return
	}

	var debts []models.Debt
	err := database.DB.Where("user_id = ?", userIDInt).Find(&debts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dívidas"})
		return
	}

	var response []models.DebtResponse
	for _, debt := range debts {
		response = append(response, models.DebtResponse{
			DebtID:       debt.DebtID,
			Value:        debt.Value,
			MaturityDate: debt.MaturityDate,
		})
	}

	c.JSON(http.StatusOK, response)
}

func PostDebts(c *gin.Context) {
	var debt models.Debt

	userRole, exists := c.Get("user_role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userRoleUint, ok := userRole.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar permissões do usuário"})
		return
	}

	if userRoleUint != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Apenas administradores podem criar dívidas"})
		return
	}

	if err := c.ShouldBindJSON(&debt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos ou incompletos"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, debt.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuário não encontrado"})
		return
	}

	debt.DebtID = uuid.New()

	if err := database.DB.Create(&debt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar dívida"})
		return
	}

	response := models.DebtResponse{
		DebtID:       debt.DebtID,
		Value:        debt.Value,
		MaturityDate: debt.MaturityDate,
	}

	c.JSON(http.StatusCreated, response)
}
