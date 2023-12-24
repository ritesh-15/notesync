package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/ritesh-15/notesync-backend/controllers"
	"github.com/ritesh-15/notesync-backend/middleware"
	"github.com/ritesh-15/notesync-backend/utils"
)

func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 60 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(utils.NewApiError("you have reached maximum request limit please try again later", nil))
		},
	}))

	auth.Post("/register", middleware.Validation(&controllers.RegisterReq{}), controllers.Register)

	auth.Post("/verify", controllers.Verify)

	auth.Post("/login", middleware.Validation(&controllers.LoginReq{}), controllers.Login)
}
