package dtos

import "github.com/ritesh-15/notesync-backend/models"

type UserDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
}

func NewUser(user *models.User) *UserDto {
	return &UserDto{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		IsVerified: user.IsVerified,
	}
}
