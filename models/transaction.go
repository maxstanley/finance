package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/satori/go.uuid"
)

type Transaction struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey"`
	StatementID uuid.UUID

	Date time.Time
	Description string
	Credit uint64
	Debit uint64
	Type string
	Essential bool
	PrimaryCategory string
	SecondaryCategory string
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewV4()

	return 
}

