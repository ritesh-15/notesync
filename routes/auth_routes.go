package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ritesh-15/notesync-backend/controllers"
)

func AuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")

	auth.POST("/register", controllers.Register)

	// auth.POST("/login")
}
