package models

import "time"

type Transaction struct {
	ID         int64     `gorm:"id"`
	MerchantID int64     `gorm:"merchant_id"`
	OutletID   int64     `gorm:"outlet_id"`
	BillTotal  float64   `gorm:"bill_total"`
	CreatedAt  time.Time `gorm:"created_at"`
	CreatedBy  int64     `gorm:"created_by"`
	UpdatedAt  time.Time `gorm:"updated_at"`
	UpdatedBy  int64     `gorm:"updated_by"`
}

func (Transaction) TableName() string {
	return "Transactions"
}
