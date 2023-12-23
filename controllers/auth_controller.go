package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/models"
	"github.com/ritesh-15/notesync-backend/utils"
)

type RegisterReq struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.NewApiError("unprocessable entity", err.Error()))
	}

	if err := config.DB.Where("email = ?", req.Email).First(&models.User{}); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewApiError("email address is already taken by another user", err.Error))
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := config.DB.Create(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewApiError("something went wrong while creating user", err.Error))
	}

	// send email
}
