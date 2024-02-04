package model

import "gorm.io/gorm"

func init() { register(Post{}) }

const (
	PostStatusDraft     = ""
	PostStatusPublish   = "publish"
	PostStatusProtected = "protected"
	PostStatusPrivate   = "private"
)

type Post struct {
	gorm.Model
	Author  uint
	Title   string
	Summary string
	Content string
	Status  string

	AllowPing    bool
	AllowComment bool
}
