package model

import "gorm.io/gorm"

func init() { register(User{}, UserMeta{}) }

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Display  string
	Password []byte
	Email    string
	Salt     string
	Avatar   string
}

const (
	UserMetaKeyRole = "role" // user role privilege
)

const (
	RoleSuperAdmin = "superAdmin"
)

type UserMeta struct {
	gorm.Model
	UserID uint
	Key    string
	Value  string
}
