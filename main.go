package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritesh-15/notesync-backend/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	app := gin.New()

	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "healt check successful ✅",
		})
	})

	app.Run(config.HOST + ":" + config.PORT)
}
