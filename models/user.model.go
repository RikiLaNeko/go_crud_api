package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username,omitempty"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email,omitempty"`
	Password  string    `gorm:"not null" json:"password,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

