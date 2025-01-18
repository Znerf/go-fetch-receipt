package models

import (

	"github.com/google/uuid"

	"gorm.io/gorm"
)


type ReceiptBase struct {
	ID 				  BINARY16 		  `gorm:"primaryKey" json:"id"`
	Retailer          string          `gorm:"type:varchar(255)" json:"retailer" validate:"required"`
	PurchaseDate      string          `gorm:"type:varchar(255)" json:"purchaseDate" validate:"required,date"`
	PurchaseTime      string          `gorm:"type:varchar(255)" json:"purchaseTime" validate:"required,time"`
	Total			  string          `gorm:"type:varchar(255)" json:"total"`
}

type Receipt struct {
	ReceiptBase
	Items []Items `gorm:"foreignKey:ReceiptId" json:"items"`
}

type Items struct {
	ID 				  	BINARY16 		  `gorm:"primaryKey" json:"id"`
	ReceiptId           *BINARY16         `json:"receiptId"`
	ShortDescription    string            `gorm:"type:varchar(255)" json:"shortDescription"`
    Price               string            `gorm:"type:varchar(255)" json:"price" validate:"float"`
}

func (m *ReceiptBase) BeforeCreate(_ *gorm.DB) error {
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

func (c ReceiptBase) TableName() string {
	return "receipts"
}

func (c Items) TableName() string {
	return "items"
}