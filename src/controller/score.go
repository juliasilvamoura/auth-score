package controller

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliasilvamoura/auth-score/src/database"
	"github.com/juliasilvamoura/auth-score/src/models"
)

func GetScore(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao interpretar o ID do usuário"})
		return
	}

	var debts []models.Debt
	if err := database.DB.Where("user_id = ?", userIDUint).Find(&debts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dívidas"})
		return
	}

	if len(debts) == 0 {
		c.JSON(http.StatusOK, gin.H{"score": 1000})
		return
	}

	var total float64
	for _, debt := range debts {
		total += float64(debt.Value)
	}
	media := total / float64(len(debts))

	score := 10000 / math.Sqrt(media+100)

	if score > 1000 {
		score = 1000
	}

	fator := math.Pow(10, float64(2))

	c.JSON(http.StatusOK, gin.H{"score": math.Trunc(score*fator) / fator})
}
