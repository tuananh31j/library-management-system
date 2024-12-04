package model

import (
	"time"

	"github.com/google/uuid"
)

type BorrowerBooks struct {
	ID         uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	BorrowerID int       `json:"borrower_id" gorm:"foreignKey:ID;references:BorrowerID"`
	BookID     int       `json:"book_id" gorm:"foreignKey:ID;references:BookID"`
	Book       *Book     `gorm:"foreignKey:BookID;references:ID"`
	Borrower   *Borrower `gorm:"foreignKey:BorrowerID;references:ID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
