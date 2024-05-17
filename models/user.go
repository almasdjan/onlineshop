package models

type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
