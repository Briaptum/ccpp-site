package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"not null" json:"title" binding:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Summary     string         `gorm:"type:text" json:"summary"`
	Date        time.Time      `gorm:"not null" json:"date" binding:"required"`
	Time        string         `json:"time"`
	Location    string         `json:"location"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

