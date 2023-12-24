package models

import (
	"time"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

type User struct {
	ID                string    `json:"id" gorm:"primaryKey"`
	Name              string    `json:"name" gorm:"notNull"`
	Email             string    `json:"email" gorm:"unique;notNull;index"`
	IsVerified        bool      `json:"isVerified"`
	Provider          string    `json:"provider"`
	VerificationToken string    `json:"verificationToken"`
	CreatedAt         time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = nanoid.New()
	return
}
