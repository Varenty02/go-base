package authmodel

import "time"

//dto vào register
type UserCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phone_no"`
}

//dto vào login
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


//entity user
type User struct {
	Id       uint    `gorm:"column:id;autoIncrement"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	PhoneNo  string `gorm:"column:phone_no"`
	RefreshToken *string `gorm:"column:refresh_token"`
	ExpiresAt *time.Time ` gorm:"column:expires_at"`
}


func (User) TableName() string       { return "users" }