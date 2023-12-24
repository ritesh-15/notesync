package models

import (
	"time"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

type Session struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Token     string    `json:"token" gorm:"unique;notNull;index"`
	UserId    string    `json:"userId" gorm:"notNull"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = nanoid.New()
	return
}
