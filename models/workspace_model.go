package models

import (
	"time"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

type Workspace struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	Title            string    `json:"title" gorm:"notNull"`
	Content          string    `json:"content"`
	Logo             string    `json:"logo"`
	BannerURL        string    `json:"bannerUrl"`
	InTrash          bool      `json:"inTrash" gorm:"default:false"`
	WorkspaceOwnerId string    `json:"workspaeOwnerId" gorm:"notNull"`
	WorkspaceOwner   string    `json:"workspaceOwner" gorm:"foreignKey:WorkspaceOwnerId; constraint:OnDelete:CASCADE"`
	CreatedAt        time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (w *Workspace) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID = nanoid.New()
	return
}
