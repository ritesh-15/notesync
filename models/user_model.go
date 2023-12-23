package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"notNull"`
	Email      string `json:"email" gorm:"unique;notNull;index"`
	IsVerified bool   `json:"isVerified"`
	Provider   string `json:"provider"`
	CreatedAt  string `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  string `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
