package model

import (
	"time"
)

//struct
type User struct {
	Id           string `json:"iduser"`
	Firstname    string `json:"firtsname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Imageprofile string `json:"imageprofile"`
	CreateAt     time.Time
	UpdateAt     time.Time
}

//slice
type Users []User

//constructor
func NewUser() *User {
	return &User{
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}
