package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ritesh-15/notesync-backend/controllers"
	"github.com/ritesh-15/notesync-backend/middleware"
)

func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/register", middleware.Validation(&controllers.RegisterReq{}), controllers.Register)
}
