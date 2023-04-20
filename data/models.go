package data

import "time"

type User struct {
	ID         int       `json:"id" sql:"id"`
	Email      string    `json:"email" validate:"required" sql:"email"`
	Password   string    `json:"password" validate:"required" sql:"password"`
	Username   string    `json:"username" validate:"required" sql:"username"`
	IsVerified bool      `json:"isverified" sql:"isverified"`
	CreatedAt  time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt  time.Time `json:"updatedat" sql:"updatedat"`
}
