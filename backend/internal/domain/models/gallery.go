package models

import (
	"time"

	"gorm.io/gorm"
)

type Gallery struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Filename  string         `gorm:"not null" json:"filename"`
	Path      string         `gorm:"not null" json:"path"`
	Category  string         `gorm:"not null" json:"category" binding:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

