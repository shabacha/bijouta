package model

import "time"

type Product struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	UserID      uint    `json:"user_id"`
	User        User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
	// Images      []string `json:"images"`
	// Categories []string   `gorm:"type:text[]" json:"categories"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Product) TableName() string { return "products" }
