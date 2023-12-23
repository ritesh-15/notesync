package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/routes"
	"github.com/ritesh-15/notesync-backend/utils"
)

func init() {
	config.LoadEnv()
	config.InitDb()
}

func main() {
	app := gin.New()

	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, utils.NewResponse(true, "health check successfull ✅", nil))
	})

	routes.InitRoute(app)

	log.Fatal(app.Run(config.HOST + ":" + config.PORT))
}
