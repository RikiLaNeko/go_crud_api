package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Username  string    `gorm:"not null" json:"username,omitempty" binding:"required"`
	Email     string    `gorm:"not null" json:"email,omitempty" binding:"required,email"`
	Password  string    `gorm:"not null" json:"password,omitempty" binding:"required"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
