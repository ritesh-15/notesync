package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/dtos"
	"github.com/ritesh-15/notesync-backend/global"
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

	url := fmt.Sprintf("%s/verify-email?token=%s&userId=%s&action=verification", config.CLIENT_URL, verificationToken, user.ID)

	// send verification email
	html := fmt.Sprintf("<a href=\"%s\">Verify</a>", url)

	go utils.SendEmail(utils.NewEmail(req.Email, url, html, "Notesync: Verification Email"))

	return c.JSON(utils.NewResponse("ok", dtos.NewUser(user)))
}

type VerifyReq struct {
	UserId string `json:"userId" validate:"required"`
	Token  string `json:"token" validate:"required"`
	Action string `json:"action" validate:"required"`
}

func Verify(c *fiber.Ctx) error {
	req := &VerifyReq{
		UserId: c.Query("userId"),
		Token:  c.Query("token"),
		Action: c.Query("action"),
	}

	if errs := global.MyValidator.Validate(req); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"validation failed on field %s, condition: %s",
				err.FailedField,
				err.Tag,
			))
		}

		return c.Status(http.StatusUnprocessableEntity).JSON(
			utils.NewApiError("unprocessable entity", errMsgs),
		)
	}

	// verify received token
	claims, err := utils.VerifyToken(req.Token, config.VERIFICATION_TOKEN_SECRET)

	if err != nil {
		return c.Status(http.StatusForbidden).JSON(utils.NewApiError("toke is not valid please try again", nil))
	}

	// check the user in database with given verification token
	var user models.User

	if result := config.DB.Where("id= ? AND verification_token= ?", claims.ID, req.Token).First(&user); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(utils.NewApiError("user not found with given id", nil))
	}

	user.VerificationToken = ""

	if !user.IsVerified || req.Action == "verification" {
		user.IsVerified = true
	}

	if result := config.DB.Save(&user); result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			utils.NewApiError("unable to save user details", nil),
		)
	}

	// generate new access and refresh tokens
	accessToken, refreshToken := utils.GenerateAccessAndRefreshToken(user.ID)

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		Secure:   false,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		Secure:   false,
	})

	return c.JSON(utils.NewResponse("ok", dtos.NewUser(&user)))
}

type LoginReq struct {
	Email string `json:"email" validate:"required,email"`
}

func Login(c *fiber.Ctx) error {
	var req LoginReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewApiError("unprocessable entity", nil))
	}

	// find user with given email
	var user models.User

	if result := config.DB.Where("email=?", req.Email).First(&user); result.Error == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).JSON(utils.NewApiError("user not found with given email address", nil))
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

	url := fmt.Sprintf("%s/verify-email?token=%s&userId=%s&action=login", config.CLIENT_URL, verificationToken, user.ID)

	// send verification email
	html := fmt.Sprintf("<a href=\"%s\">Verify</a>", url)

	go utils.SendEmail(utils.NewEmail(req.Email, url, html, "Notesync: Login verification email"))

	return c.JSON(utils.NewResponse("ok", nil))
}
