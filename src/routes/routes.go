package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juliasilvamoura/auth-score/src/controller"
)

func HandleRequests(r *gin.Engine, auth *gin.RouterGroup) {

	// Login
	auth.POST("/login", controller.Login)
	// Logout
	auth.POST("/logout", controller.Logout)
	// Create User
	r.POST("/register", controller.PostUser)
	// Get debts
	auth.GET("/debts", controller.GetAllDebts)
	// Create debts
	auth.POST("/debts", controller.PostDebts)
	// Score
	auth.GET("/score", controller.GetScore)

}
