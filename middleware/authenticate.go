package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/models"
	"github.com/ritesh-15/notesync-backend/utils"
)

func Authenticate(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")

	log.Info(accessToken)

	if accessToken == "" {
		return c.Status(http.StatusUnauthorized).JSON(utils.NewApiError("jwt token not found", nil))
	}

	// verify jwt token
	claims, err := utils.VerifyToken(accessToken, config.ACCESS_TOKEN_SECRET)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(utils.NewApiError("you are not authorized to access this resource", nil))
	}

	// find the user with id
	var user models.User
	if result := config.DB.Where("id = ?", claims.ID).First(&user); result.Error != nil {
		return c.Status(http.StatusUnauthorized).JSON(utils.NewApiError("you are not authorized to access this resource", nil))
	}

	c.Locals("user", user)
	return c.Next()
}
