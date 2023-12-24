package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/models"
	"github.com/ritesh-15/notesync-backend/utils"
	"gorm.io/gorm"
)

type RegisterReq struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	var req RegisterReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewApiError("unprocessable entity", nil))
	}

	if result := config.DB.Where("email=?", req.Email).First(&models.User{}); result.Error != gorm.ErrRecordNotFound || result.RowsAffected > 0 {
		return c.Status(http.StatusBadRequest).JSON(utils.NewApiError("email address is already taken by another user", nil))
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if result := config.DB.Create(user); result.Error != nil {
		log.Info(result)
		return c.Status(http.StatusBadRequest).JSON(utils.NewApiError("error while creating user", nil))
	}

	// generate verification token
	verificationToken, err := utils.SignVerificationToken(user.ID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.NewApiError("error while signing verification token", nil))
	}

	user.VerificationToken = verificationToken

	if result := config.DB.Save(user); result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.NewApiError("error while saving verification token", nil))
	}

	// send verification email

	return c.JSON(utils.NewResponse("ok", user))
}
