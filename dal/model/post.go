package model

import "gorm.io/gorm"

func init() { register(Post{}) }

const (
	PostStatusDraft     = ""
	PostStatusPublish   = "publish"
	PostStatusProtected = "protected"
	PostStatusPrivate   = "private"
)
const (
	BitPing = 1 << iota
	BitComment
)

type Post struct {
	gorm.Model
	AuthorID uint `gorm:"index"`
	Author   User `gorm:"foreignKey:author_id"`
	Title    string
	Summary  string
	Content  string
	Status   string
	Disallow int
}

func (p *Post) set(b int, allow bool) {
	if allow {
		p.Disallow = p.Disallow &^ b
	} else {
		p.Disallow = p.Disallow | b
	}
}

func (p *Post) allow(flag int) bool {
	return p.Disallow&flag == 0
}
func (p *Post) SetAllowPing(b bool) {
	p.set(BitPing, b)
}
func (p *Post) SetAllowComment(b bool) {
	p.set(BitComment, b)
}
func (p *Post) AllowPing() bool {
	return p.allow(BitPing)
}
func (p *Post) AllowComment() bool {
	return p.allow(BitComment)
}
