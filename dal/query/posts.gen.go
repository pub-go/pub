// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"code.gopub.tech/pub/dal/model"
)

func newPost(db *gorm.DB, opts ...gen.DOOption) post {
	_post := post{}

	_post.postDo.UseDB(db, opts...)
	_post.postDo.UseModel(&model.Post{})

	tableName := _post.postDo.TableName()
	_post.ALL = field.NewAsterisk(tableName)
	_post.ID = field.NewUint(tableName, "id")
	_post.CreatedAt = field.NewTime(tableName, "created_at")
	_post.UpdatedAt = field.NewTime(tableName, "updated_at")
	_post.DeletedAt = field.NewField(tableName, "deleted_at")
	_post.AuthorID = field.NewUint(tableName, "author_id")
	_post.Title = field.NewString(tableName, "title")
	_post.Summary = field.NewString(tableName, "summary")
	_post.Content = field.NewString(tableName, "content")
	_post.Status = field.NewString(tableName, "status")
	_post.Disallow = field.NewInt(tableName, "disallow")
	_post.Author = postBelongsToAuthor{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Author", "model.User"),
	}

	_post.fillFieldMap()

	return _post
}

type post struct {
	postDo postDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	AuthorID  field.Uint
	Title     field.String
	Summary   field.String
	Content   field.String
	Status    field.String
	Disallow  field.Int
	Author    postBelongsToAuthor

	fieldMap map[string]field.Expr
}

func (p post) Table(newTableName string) *post {
	p.postDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p post) As(alias string) *post {
	p.postDo.DO = *(p.postDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *post) updateTableName(table string) *post {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.AuthorID = field.NewUint(table, "author_id")
	p.Title = field.NewString(table, "title")
	p.Summary = field.NewString(table, "summary")
	p.Content = field.NewString(table, "content")
	p.Status = field.NewString(table, "status")
	p.Disallow = field.NewInt(table, "disallow")

	p.fillFieldMap()

	return p
}

func (p *post) WithContext(ctx context.Context) *postDo { return p.postDo.WithContext(ctx) }

func (p post) TableName() string { return p.postDo.TableName() }

func (p post) Alias() string { return p.postDo.Alias() }

func (p post) Columns(cols ...field.Expr) gen.Columns { return p.postDo.Columns(cols...) }

func (p *post) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *post) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["author_id"] = p.AuthorID
	p.fieldMap["title"] = p.Title
	p.fieldMap["summary"] = p.Summary
	p.fieldMap["content"] = p.Content
	p.fieldMap["status"] = p.Status
	p.fieldMap["disallow"] = p.Disallow

}

func (p post) clone(db *gorm.DB) post {
	p.postDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p post) replaceDB(db *gorm.DB) post {
	p.postDo.ReplaceDB(db)
	return p
}

type postBelongsToAuthor struct {
	db *gorm.DB

	field.RelationField
}

func (a postBelongsToAuthor) Where(conds ...field.Expr) *postBelongsToAuthor {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a postBelongsToAuthor) WithContext(ctx context.Context) *postBelongsToAuthor {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a postBelongsToAuthor) Session(session *gorm.Session) *postBelongsToAuthor {
	a.db = a.db.Session(session)
	return &a
}

func (a postBelongsToAuthor) Model(m *model.Post) *postBelongsToAuthorTx {
	return &postBelongsToAuthorTx{a.db.Model(m).Association(a.Name())}
}

type postBelongsToAuthorTx struct{ tx *gorm.Association }

func (a postBelongsToAuthorTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a postBelongsToAuthorTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a postBelongsToAuthorTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a postBelongsToAuthorTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a postBelongsToAuthorTx) Clear() error {
	return a.tx.Clear()
}

func (a postBelongsToAuthorTx) Count() int64 {
	return a.tx.Count()
}

type postDo struct{ gen.DO }

func (p postDo) Debug() *postDo {
	return p.withDO(p.DO.Debug())
}

func (p postDo) WithContext(ctx context.Context) *postDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postDo) ReadDB() *postDo {
	return p.Clauses(dbresolver.Read)
}

func (p postDo) WriteDB() *postDo {
	return p.Clauses(dbresolver.Write)
}

func (p postDo) Session(config *gorm.Session) *postDo {
	return p.withDO(p.DO.Session(config))
}

func (p postDo) Clauses(conds ...clause.Expression) *postDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postDo) Returning(value interface{}, columns ...string) *postDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postDo) Not(conds ...gen.Condition) *postDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postDo) Or(conds ...gen.Condition) *postDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postDo) Select(conds ...field.Expr) *postDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postDo) Where(conds ...gen.Condition) *postDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postDo) Order(conds ...field.Expr) *postDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postDo) Distinct(cols ...field.Expr) *postDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postDo) Omit(cols ...field.Expr) *postDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postDo) Join(table schema.Tabler, on ...field.Expr) *postDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postDo) LeftJoin(table schema.Tabler, on ...field.Expr) *postDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postDo) RightJoin(table schema.Tabler, on ...field.Expr) *postDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postDo) Group(cols ...field.Expr) *postDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postDo) Having(conds ...gen.Condition) *postDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postDo) Limit(limit int) *postDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postDo) Offset(offset int) *postDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *postDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postDo) Unscoped() *postDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postDo) Create(values ...*model.Post) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postDo) CreateInBatches(values []*model.Post, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postDo) Save(values ...*model.Post) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postDo) First() (*model.Post, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) Take() (*model.Post, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) Last() (*model.Post, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) Find() ([]*model.Post, error) {
	result, err := p.DO.Find()
	return result.([]*model.Post), err
}

func (p postDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Post, err error) {
	buf := make([]*model.Post, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postDo) FindInBatches(result *[]*model.Post, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postDo) Attrs(attrs ...field.AssignExpr) *postDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postDo) Assign(attrs ...field.AssignExpr) *postDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postDo) Joins(fields ...field.RelationField) *postDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postDo) Preload(fields ...field.RelationField) *postDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postDo) FirstOrInit() (*model.Post, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) FirstOrCreate() (*model.Post, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) FindByPage(offset int, limit int) (result []*model.Post, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postDo) Delete(models ...*model.Post) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postDo) withDO(do gen.Dao) *postDo {
	p.DO = *do.(*gen.DO)
	return p
}
