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

func newUserMeta(db *gorm.DB, opts ...gen.DOOption) userMeta {
	_userMeta := userMeta{}

	_userMeta.userMetaDo.UseDB(db, opts...)
	_userMeta.userMetaDo.UseModel(&model.UserMeta{})

	tableName := _userMeta.userMetaDo.TableName()
	_userMeta.ALL = field.NewAsterisk(tableName)
	_userMeta.ID = field.NewUint(tableName, "id")
	_userMeta.CreatedAt = field.NewTime(tableName, "created_at")
	_userMeta.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userMeta.DeletedAt = field.NewField(tableName, "deleted_at")
	_userMeta.UserID = field.NewUint(tableName, "user_id")
	_userMeta.Key = field.NewString(tableName, "key")
	_userMeta.Value = field.NewString(tableName, "value")

	_userMeta.fillFieldMap()

	return _userMeta
}

type userMeta struct {
	userMetaDo userMetaDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserID    field.Uint
	Key       field.String
	Value     field.String

	fieldMap map[string]field.Expr
}

func (u userMeta) Table(newTableName string) *userMeta {
	u.userMetaDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userMeta) As(alias string) *userMeta {
	u.userMetaDo.DO = *(u.userMetaDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userMeta) updateTableName(table string) *userMeta {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.UserID = field.NewUint(table, "user_id")
	u.Key = field.NewString(table, "key")
	u.Value = field.NewString(table, "value")

	u.fillFieldMap()

	return u
}

func (u *userMeta) WithContext(ctx context.Context) *userMetaDo { return u.userMetaDo.WithContext(ctx) }

func (u userMeta) TableName() string { return u.userMetaDo.TableName() }

func (u userMeta) Alias() string { return u.userMetaDo.Alias() }

func (u userMeta) Columns(cols ...field.Expr) gen.Columns { return u.userMetaDo.Columns(cols...) }

func (u *userMeta) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userMeta) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 7)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["key"] = u.Key
	u.fieldMap["value"] = u.Value
}

func (u userMeta) clone(db *gorm.DB) userMeta {
	u.userMetaDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userMeta) replaceDB(db *gorm.DB) userMeta {
	u.userMetaDo.ReplaceDB(db)
	return u
}

type userMetaDo struct{ gen.DO }

func (u userMetaDo) Debug() *userMetaDo {
	return u.withDO(u.DO.Debug())
}

func (u userMetaDo) WithContext(ctx context.Context) *userMetaDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userMetaDo) ReadDB() *userMetaDo {
	return u.Clauses(dbresolver.Read)
}

func (u userMetaDo) WriteDB() *userMetaDo {
	return u.Clauses(dbresolver.Write)
}

func (u userMetaDo) Session(config *gorm.Session) *userMetaDo {
	return u.withDO(u.DO.Session(config))
}

func (u userMetaDo) Clauses(conds ...clause.Expression) *userMetaDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userMetaDo) Returning(value interface{}, columns ...string) *userMetaDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userMetaDo) Not(conds ...gen.Condition) *userMetaDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userMetaDo) Or(conds ...gen.Condition) *userMetaDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userMetaDo) Select(conds ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userMetaDo) Where(conds ...gen.Condition) *userMetaDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userMetaDo) Order(conds ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userMetaDo) Distinct(cols ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userMetaDo) Omit(cols ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userMetaDo) Join(table schema.Tabler, on ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userMetaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userMetaDo) RightJoin(table schema.Tabler, on ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userMetaDo) Group(cols ...field.Expr) *userMetaDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userMetaDo) Having(conds ...gen.Condition) *userMetaDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userMetaDo) Limit(limit int) *userMetaDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userMetaDo) Offset(offset int) *userMetaDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userMetaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userMetaDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userMetaDo) Unscoped() *userMetaDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userMetaDo) Create(values ...*model.UserMeta) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userMetaDo) CreateInBatches(values []*model.UserMeta, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userMetaDo) Save(values ...*model.UserMeta) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userMetaDo) First() (*model.UserMeta, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserMeta), nil
	}
}

func (u userMetaDo) Take() (*model.UserMeta, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserMeta), nil
	}
}

func (u userMetaDo) Last() (*model.UserMeta, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserMeta), nil
	}
}

func (u userMetaDo) Find() ([]*model.UserMeta, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserMeta), err
}

func (u userMetaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserMeta, err error) {
	buf := make([]*model.UserMeta, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userMetaDo) FindInBatches(result *[]*model.UserMeta, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userMetaDo) Attrs(attrs ...field.AssignExpr) *userMetaDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userMetaDo) Assign(attrs ...field.AssignExpr) *userMetaDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userMetaDo) Joins(fields ...field.RelationField) *userMetaDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userMetaDo) Preload(fields ...field.RelationField) *userMetaDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userMetaDo) FirstOrInit() (*model.UserMeta, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserMeta), nil
	}
}

func (u userMetaDo) FirstOrCreate() (*model.UserMeta, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserMeta), nil
	}
}

func (u userMetaDo) FindByPage(offset int, limit int) (result []*model.UserMeta, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userMetaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userMetaDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userMetaDo) Delete(models ...*model.UserMeta) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userMetaDo) withDO(do gen.Dao) *userMetaDo {
	u.DO = *do.(*gen.DO)
	return u
}
