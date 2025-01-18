package models

import (
	// "frail-check-api/constants"
	// "time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// Base contains common columns for all tables.
// type Base struct{
	
// }
type Receipt2 struct {
	ID 				  BINARY16 		  `gorm:"primaryKey" json:"id"`
	Retailer          string          `gorm:"type:varchar(255)" json:"retailer" validate:"required"`
	PurchaseDate      string          `gorm:"type:varchar(255)" json:"purchaseDate" validate:"required"`
	PurchaseTime      string          `gorm:"type:varchar(255)" json:"purchaseTime" validate:"required"`
	Total			  string          `gorm:"type:varchar(255)" json:"total"`
}

type Receipt struct {
	Receipt2
	Items []Items `gorm:"foreignKey:ReceiptId" json:"items"`
}

type Items struct {
	ID 				  	BINARY16 		  `gorm:"primaryKey" json:"id"`
	ReceiptId           *BINARY16         `json:"receiptId"`
	ShortDescription    string            `gorm:"type:varchar(255)" json:"shortDescription"`
    Price               string            `gorm:"type:varchar(255)" json:"price"`
}

func (m *Receipt2) BeforeCreate(_ *gorm.DB) error {
	defaultId, _ := StringToBinary16("00000000-0000-0000-0000-000000000000")
	if m.ID == defaultId {
		id, err := uuid.NewRandom()
		m.ID = BINARY16(id)
		return err
	}
	return nil
}

func (m *Items) BeforeCreate(_ *gorm.DB) error {
	defaultId, _ := StringToBinary16("00000000-0000-0000-0000-000000000000")
	if m.ID == defaultId {
		id, err := uuid.NewRandom()
		m.ID = BINARY16(id)
		return err
	}
	return nil
}

func (c Receipt2) TableName() string {
	return "receipts"
}

func (c Items) TableName() string {
	return "items"
}