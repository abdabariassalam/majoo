package models

import "time"

type Merchant struct {
	ID           int64     `gorm:"id"`
	UserID       int64     `gorm:"user_id"`
	MerchantName string    `gorm:"merchant_name"`
	CreatedAt    time.Time `gorm:"created_at"`
	CreatedBy    int64     `gorm:"created_by"`
	UpdatedAt    time.Time `gorm:"updated_at"`
	UpdatedBy    int64     `gorm:"updated_by"`
}

func (Merchant) TableName() string {
	return "Merchants"
}

type MerchantMonthReport struct {
	Date         string `gorm:"tgl"`
	MerchantName string `gorm:"merchant_name"`
	Omzet        int64  `gorm:"omzet"`
}
