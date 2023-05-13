package model

import "time"

type Category struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Category) TableName() string { return "categories" }
