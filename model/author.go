package model

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at"`
}
