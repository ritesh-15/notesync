package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/routes"
	"github.com/ritesh-15/notesync-backend/utils"
)

func init() {
	config.LoadEnv()
	config.InitDb()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(
				utils.NewApiError("something went wrong at our side please try again later", nil),
			)
		},
	})

	app.Use(helmet.New())
	app.Use(logger.New())

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(utils.NewResponse("health check successfull ✅", nil))
	})

	routes.InitRoute(app)

	log.Fatal(app.Listen(config.HOST + ":" + config.PORT))
}
