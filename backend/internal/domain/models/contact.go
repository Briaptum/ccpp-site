package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	FirstName string         `gorm:"not null" json:"firstName" binding:"required"`
	LastName  string         `gorm:"not null" json:"lastName" binding:"required"`
	Email     string         `gorm:"not null" json:"email" binding:"required,email"`
	Phone     string         `json:"phone"`
	Subject   string         `gorm:"not null" json:"subject" binding:"required"`
	Message   string         `gorm:"type:text;not null" json:"message" binding:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

