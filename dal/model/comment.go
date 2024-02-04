package model

import "gorm.io/gorm"

func init() { register(Comment{}) }

const (
	CommentStatusWaiting = "waiting" // 等待审核
	CommentStatusApprove = "approve" // 审核通过可展示
)

type Comment struct {
	gorm.Model
	AuthorID    uint   // 评论人如果是注册用户
	AuthorFid   string // FakeID 游客随机标记 等待审核时展示自己的
	AuthorName  string
	AuthorEmail string
	AuthorURL   string

	Parent uint // 回复哪条评论
	Depth  uint // 回复时嵌套层数

	PostID  uint   // 所属文章
	Content string // 评论内容
}
