// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q        = new(Query)
	Comment  *comment
	Option   *option
	Post     *post
	User     *user
	UserMeta *userMeta
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Comment = &Q.Comment
	Option = &Q.Option
	Post = &Q.Post
	User = &Q.User
	UserMeta = &Q.UserMeta
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:       db,
		Comment:  newComment(db, opts...),
		Option:   newOption(db, opts...),
		Post:     newPost(db, opts...),
		User:     newUser(db, opts...),
		UserMeta: newUserMeta(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Comment  comment
	Option   option
	Post     post
	User     user
	UserMeta userMeta
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:       db,
		Comment:  q.Comment.clone(db),
		Option:   q.Option.clone(db),
		Post:     q.Post.clone(db),
		User:     q.User.clone(db),
		UserMeta: q.UserMeta.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:       db,
		Comment:  q.Comment.replaceDB(db),
		Option:   q.Option.replaceDB(db),
		Post:     q.Post.replaceDB(db),
		User:     q.User.replaceDB(db),
		UserMeta: q.UserMeta.replaceDB(db),
	}
}

type queryCtx struct {
	Comment  *commentDo
	Option   *optionDo
	Post     *postDo
	User     *userDo
	UserMeta *userMetaDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Comment:  q.Comment.WithContext(ctx),
		Option:   q.Option.WithContext(ctx),
		Post:     q.Post.WithContext(ctx),
		User:     q.User.WithContext(ctx),
		UserMeta: q.UserMeta.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
