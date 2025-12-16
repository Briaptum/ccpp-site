package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	Name         string         `json:"name" gorm:"not null"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-"` // stored hash, never serialized to clients
	Role         string         `json:"role" gorm:"not null;default:user"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Role     *string `json:"role,omitempty"`
}
