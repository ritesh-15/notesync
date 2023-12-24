package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/dtos"
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
	var req VerifyReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewApiError("unprocessable entity", nil))
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

	session := models.Session{
		UserId: user.ID,
		Token:  refreshToken,
	}

	if result := config.DB.Create(&session); result.Error != nil {
		return fiber.ErrInternalServerError
	}

	c.Cookie(utils.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*24*30)))

	c.Cookie(utils.SetCookie(c, "refresh_token", refreshToken, time.Now().Add(time.Hour*24*30)))

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

type RefreshReq struct {
	AccessToken  string `cookie:"access_token" validate:"required"`
	RefreshToken string `cookie:"refresh_token" validate:"required"`
}

func RefreshTokens(c *fiber.Ctx) error {
	var tokens RefreshReq

	if err := c.CookieParser(&tokens); err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.NewApiError("tokens not found", nil))
	}

	// verify refresh token
	claim, err := utils.VerifyToken(tokens.RefreshToken, config.REFRESH_TOKEN_SECRET)

	if err != nil {

		// revoke the session
		if result := config.DB.Where("token = ?", tokens.RefreshToken).Delete(&models.Session{}); result.Error != nil {
			log.Error(result.Error)
			return fiber.ErrInternalServerError
		}

		c.Cookie(utils.ClearCookie(c, "access_token"))
		c.Cookie(utils.ClearCookie(c, "refresh_token"))

		return c.Status(http.StatusUnauthorized).JSON(utils.NewApiError("token is not valid", nil))
	}

	// find session in database
	var session models.Session

	if result := config.DB.Where("user_id = ? AND token = ?", claim.ID, tokens.RefreshToken).First(&session); result.Error == gorm.ErrRecordNotFound || result.Error != nil {

		log.Error(result.Error)

		// revoke all the sessions
		if result := config.DB.Where("user_id = ?", claim.ID).Delete(&models.Session{}); result.Error != nil {
			log.Error(result.Error)
			return fiber.ErrInternalServerError
		}

		c.Cookie(utils.ClearCookie(c, "access_token"))
		c.Cookie(utils.ClearCookie(c, "refresh_token"))

		return c.Status(http.StatusUnauthorized).JSON(utils.NewApiError("session not found", nil))
	}

	// revoke previous session
	if result := config.DB.Where("token = ?", tokens.RefreshToken).Delete(&models.Session{}); result.Error != nil {
		log.Error(result.Error)
		return fiber.ErrInternalServerError
	}

	// generate new access and refresh tokens
	accessToken, refreshToken := utils.GenerateAccessAndRefreshToken(session.UserId)

	newSession := models.Session{
		UserId: session.UserId,
		Token:  refreshToken,
	}

	if result := config.DB.Create(&newSession); result.Error != nil {
		log.Error(result.Error)
		return fiber.ErrInternalServerError
	}

	c.Cookie(utils.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*24*30)))

	c.Cookie(utils.SetCookie(c, "refresh_token", refreshToken, time.Now().Add(time.Hour*24*30)))

	return c.JSON(utils.NewResponse("tokens refresh successfully!", nil))
}

func LoggedInUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	return c.Status(http.StatusOK).JSON(utils.NewResponse("user fetched successfully", dtos.NewUser(&user)))
}

func Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	c.Cookie(utils.ClearCookie(c, "access_token"))
	c.Cookie(utils.ClearCookie(c, "refresh_token"))

	// delete all the active sessions
	if result := config.DB.Where("user_id = ?", user.ID).Delete(&models.Session{}); result.Error != nil {
		log.Error(result.Error)
		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(utils.NewResponse("logged out succesfully", nil))
}
