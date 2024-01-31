package model

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Value string
}

const (
	OptionNameInstalled = "installed"
	OptionValueYes      = "1"
)
