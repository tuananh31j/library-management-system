package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Name      string    `json:"name"`
	AuthorID  uuid.UUID `json:"author_id"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	Author    *Author   `gorm:"foreignKey:AuthorID;references:ID" json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
