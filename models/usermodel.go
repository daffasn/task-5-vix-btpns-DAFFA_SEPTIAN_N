package models

import "time"

type User struct {
	ID         int `gorm:"primaryKey"`
	Username   string
	Email      string
	Password   string
	Photo      string
	Created_At time.Time `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	Updated_at time.Time `json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type Register struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyProfile struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Photo    string
}
