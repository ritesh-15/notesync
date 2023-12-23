package routes

import "github.com/gin-gonic/gin"

func InitRoute(app *gin.Engine) {
	api := app.Group("/api")

	AuthRoutes(api)
}
