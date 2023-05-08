package model

import "time"

type User struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	Username    string     `gorm:"unique" json:"username"`
	Email       string     `gorm:"unique;not null;validate:required,email"`
	Age         int        `json:"age"`
	Password    string     `json:"password"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	ConfirmedAt *time.Time `json:"confirmed_at"`
}

func (User) TableName() string { return "users" }
