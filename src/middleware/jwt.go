package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/juliasilvamoura/auth-score/src/database"
	"github.com/juliasilvamoura/auth-score/src/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt_key, exists := c.Get("JWT_SECRET")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET não encontrado"})
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verifica se o token está na blacklist
		if database.IsTokenBlacklisted(tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou revogado"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		claims, ok := token.Claims.(*models.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não foi possível interpretar as claims"})
			return
		}

		// Verifica se o usuário ainda tem o mesmo role
		var user models.User
		if err := database.DB.First(&user, claims.UserID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
			return
		}

		if user.RoleID != claims.RoleID {
			// Se o role mudou, invalida o token
			database.AddToBlacklist(tokenString, jwt.NewNumericDate(claims.ExpiresAt.Time).Sub(time.Now()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Permissões do usuário foram alteradas. Por favor, faça login novamente"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.RoleID)
		c.Next()
	}
}
