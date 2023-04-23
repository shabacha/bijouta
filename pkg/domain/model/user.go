package model

import "time"

type User struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Age         string     `json:"age"`
	Password    string     `json:"password"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	ConfirmedAt *time.Time `json:"confirmed_at"`
	Products    []Product  `gorm:"ForeignKey:UserID" json:"products"`
}

func (User) TableName() string { return "users" }
