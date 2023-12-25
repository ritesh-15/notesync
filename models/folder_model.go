package models

import (
	"time"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

type Folder struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"notNull"`
	Content     string    `json:"content"`
	Logo        string    `json:"logo"`
	BannerURL   string    `json:"bannerUrl"`
	InTrash     bool      `json:"inTrash" gorm:"default:false"`
	WorkspaceId string    `json:"workspaceId" gorm:"notNull"`
	Workspace   string    `json:"workspace" gorm:"foreignKey:WorkspaceId; constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (f *Folder) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = nanoid.New()
	return
}
