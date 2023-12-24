package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	api := app.Group("/api")
	AuthRoutes(api)
}
