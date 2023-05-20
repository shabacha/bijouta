package model

import "time"

type Cart struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Product   []Product  `json:"product"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Cart) TableName() string { return "carts" }
