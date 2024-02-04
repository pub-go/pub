package model

import "gorm.io/gorm"

func init() { register(User{}) }

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password []byte
	Email    string
	Salt     string
}
