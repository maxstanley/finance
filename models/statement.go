package models

import (
	"gorm.io/gorm"

	"github.com/satori/go.uuid"
)

type Statement struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey"`

	AccountNumber int
	SortCode int
	Transactions []Transaction `gorm:"foreignKey:StatementID"`
}

func (s *Statement) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewV4()

	return 
}

