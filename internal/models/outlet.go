package models

import "time"

type Outlet struct {
	ID         int64     `gorm:"id"`
	MerchantID int64     `gorm:"merchant_id"`
	OutletName string    `gorm:"outlet_name"`
	CreatedAt  time.Time `gorm:"created_at"`
	CreatedBy  int64     `gorm:"created_by"`
	UpdatedAt  time.Time `gorm:"updated_at"`
	UpdatedBy  int64     `gorm:"updated_by"`
}

func (Outlet) TableName() string {
	return "Outlets"
}

type OutletMonthReport struct {
	Date         string `gorm:"tgl"`
	MerchantName string `gorm:"merchant_name"`
	OtletName    string `gorm:"outlet_name"`
	Omzet        int64  `gorm:"omzet"`
}
