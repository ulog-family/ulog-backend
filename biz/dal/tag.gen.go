// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"ulog-backend/biz/model"
)

func newTag(db *gorm.DB, opts ...gen.DOOption) tag {
	_tag := tag{}

	_tag.tagDo.UseDB(db, opts...)
	_tag.tagDo.UseModel(&model.Tag{})

	tableName := _tag.tagDo.TableName()
	_tag.ALL = field.NewAsterisk(tableName)
	_tag.ID = field.NewInt64(tableName, "id")
	_tag.Name = field.NewString(tableName, "name")
	_tag.Articles = tagManyToManyArticles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Articles", "model.Article"),
	}

	_tag.fillFieldMap()

	return _tag
}

type tag struct {
	tagDo

	ALL      field.Asterisk
	ID       field.Int64
	Name     field.String
	Articles tagManyToManyArticles

	fieldMap map[string]field.Expr
}

func (t tag) Table(newTableName string) *tag {
	t.tagDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tag) As(alias string) *tag {
	t.tagDo.DO = *(t.tagDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tag) updateTableName(table string) *tag {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Name = field.NewString(table, "name")

	t.fillFieldMap()

	return t
}

func (t *tag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tag) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name

}

func (t tag) clone(db *gorm.DB) tag {
	t.tagDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tag) replaceDB(db *gorm.DB) tag {
	t.tagDo.ReplaceDB(db)
	return t
}

type tagManyToManyArticles struct {
	db *gorm.DB

	field.RelationField
}

func (a tagManyToManyArticles) Where(conds ...field.Expr) *tagManyToManyArticles {
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

func (a tagManyToManyArticles) WithContext(ctx context.Context) *tagManyToManyArticles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a tagManyToManyArticles) Session(session *gorm.Session) *tagManyToManyArticles {
	a.db = a.db.Session(session)
	return &a
}

func (a tagManyToManyArticles) Model(m *model.Tag) *tagManyToManyArticlesTx {
	return &tagManyToManyArticlesTx{a.db.Model(m).Association(a.Name())}
}

type tagManyToManyArticlesTx struct{ tx *gorm.Association }

func (a tagManyToManyArticlesTx) Find() (result []*model.Article, err error) {
	return result, a.tx.Find(&result)
}

func (a tagManyToManyArticlesTx) Append(values ...*model.Article) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a tagManyToManyArticlesTx) Replace(values ...*model.Article) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a tagManyToManyArticlesTx) Delete(values ...*model.Article) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a tagManyToManyArticlesTx) Clear() error {
	return a.tx.Clear()
}

func (a tagManyToManyArticlesTx) Count() int64 {
	return a.tx.Count()
}

type tagDo struct{ gen.DO }

func (t tagDo) Debug() *tagDo {
	return t.withDO(t.DO.Debug())
}

func (t tagDo) WithContext(ctx context.Context) *tagDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tagDo) ReadDB() *tagDo {
	return t.Clauses(dbresolver.Read)
}

func (t tagDo) WriteDB() *tagDo {
	return t.Clauses(dbresolver.Write)
}

func (t tagDo) Session(config *gorm.Session) *tagDo {
	return t.withDO(t.DO.Session(config))
}

func (t tagDo) Clauses(conds ...clause.Expression) *tagDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tagDo) Returning(value interface{}, columns ...string) *tagDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tagDo) Not(conds ...gen.Condition) *tagDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tagDo) Or(conds ...gen.Condition) *tagDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tagDo) Select(conds ...field.Expr) *tagDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tagDo) Where(conds ...gen.Condition) *tagDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tagDo) Order(conds ...field.Expr) *tagDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tagDo) Distinct(cols ...field.Expr) *tagDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tagDo) Omit(cols ...field.Expr) *tagDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tagDo) Join(table schema.Tabler, on ...field.Expr) *tagDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tagDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tagDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tagDo) RightJoin(table schema.Tabler, on ...field.Expr) *tagDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tagDo) Group(cols ...field.Expr) *tagDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tagDo) Having(conds ...gen.Condition) *tagDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tagDo) Limit(limit int) *tagDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tagDo) Offset(offset int) *tagDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tagDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tagDo) Unscoped() *tagDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tagDo) Create(values ...*model.Tag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tagDo) CreateInBatches(values []*model.Tag, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tagDo) Save(values ...*model.Tag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tagDo) First() (*model.Tag, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tag), nil
	}
}

func (t tagDo) Take() (*model.Tag, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tag), nil
	}
}

func (t tagDo) Last() (*model.Tag, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tag), nil
	}
}

func (t tagDo) Find() ([]*model.Tag, error) {
	result, err := t.DO.Find()
	return result.([]*model.Tag), err
}

func (t tagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Tag, err error) {
	buf := make([]*model.Tag, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tagDo) FindInBatches(result *[]*model.Tag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tagDo) Attrs(attrs ...field.AssignExpr) *tagDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tagDo) Assign(attrs ...field.AssignExpr) *tagDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tagDo) Joins(fields ...field.RelationField) *tagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tagDo) Preload(fields ...field.RelationField) *tagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tagDo) FirstOrInit() (*model.Tag, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tag), nil
	}
}

func (t tagDo) FirstOrCreate() (*model.Tag, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tag), nil
	}
}

func (t tagDo) FindByPage(offset int, limit int) (result []*model.Tag, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tagDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tagDo) Delete(models ...*model.Tag) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tagDo) withDO(do gen.Dao) *tagDo {
	t.DO = *do.(*gen.DO)
	return t
}