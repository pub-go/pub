package model

import "gorm.io/gorm"

func init() { register(Option{}) }

const (
	OptionNameInstalled = "installed"
	OptionValueYes      = "1"
	OptionNameSalt      = "salt"
	OptionNameSiteTitle = "site_title"
)

type Option struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Value string
}
