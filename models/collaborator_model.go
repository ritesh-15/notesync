package models

import (
	"time"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

type Collaborator struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	WorkspaceId string    `json:"workspaceId" gorm:"notNull"`
	Workspace   string    `json:"workspace" gorm:"foreignKey:WorkspaceId; constraint:OnDelete:CASCADE"`
	UserId      string    `json:"userId" gorm:"notNull"`
	User        User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (c *Collaborator) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = nanoid.New()
	return
}
