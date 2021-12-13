package models

import "time"

type User struct {
	ID        int64     `gorm:"id"`
	Name      string    `gorm:"name"`
	UserName  string    `gorm:"user_name"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
	CreatedBy int64     `gorm:"created_by"`
	UpdatedAt time.Time `gorm:"updated_at"`
	UpdatedBy int64     `gorm:"updated_by"`
}

func (User) TableName() string {
	return "Users"
}
