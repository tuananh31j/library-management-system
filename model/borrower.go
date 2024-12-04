package model

import (
	"time"

	"github.com/google/uuid"
)

type Borrower struct {
	ID        uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
