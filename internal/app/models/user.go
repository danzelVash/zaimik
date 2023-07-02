package models

type User struct {
	Id          int    `json:"-" db:"id"`
	LastName    string `json:"last_name" binding:"required" db:"last_name"`
	FirstName   string `json:"first_name" binding:"required" db:"first_name"`
	SurName     string `json:"sur_name" binding:"required" db:"first_name"`
	Email       string `json:"email" binding:"required" db:"email"`
	PhoneNumber string `json:"phone_number" binding:"required" db:"phone_number"`
	City        string `json:"city" binding:"required" db:"city"`
	Admin       bool   `json:"-" db:"is_admin"`
}
