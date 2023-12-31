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

func newArticle(db *gorm.DB, opts ...gen.DOOption) article {
	_article := article{}

	_article.articleDo.UseDB(db, opts...)
	_article.articleDo.UseModel(&model.Article{})

	tableName := _article.articleDo.TableName()
	_article.ALL = field.NewAsterisk(tableName)
	_article.ID = field.NewInt64(tableName, "id")
	_article.Title = field.NewString(tableName, "title")
	_article.UpdatedAt = field.NewTime(tableName, "updated_at")
	_article.CreatedAt = field.NewTime(tableName, "created_at")
	_article.Content = field.NewString(tableName, "content")
	_article.Like = field.NewInt64(tableName, "like")
	_article.Read = field.NewInt64(tableName, "read")
	_article.Category = field.NewString(tableName, "category")
	_article.Password = field.NewString(tableName, "password")
	_article.Tags = articleManyToManyTags{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Tags", "model.Tag"),
		Articles: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Tags.Articles", "model.Article"),
		},
	}

	_article.Authors = articleManyToManyAuthors{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Authors", "model.User"),
		Articles: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Authors.Articles", "model.Article"),
		},
		Likes: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Authors.Likes", "model.Article"),
		},
	}

	_article.Fans = articleManyToManyFans{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Fans", "model.User"),
		Articles: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Fans.Articles", "model.Article"),
		},
		Likes: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Fans.Likes", "model.Article"),
		},
	}

	_article.fillFieldMap()

	return _article
}

type article struct {
	articleDo

	ALL       field.Asterisk
	ID        field.Int64
	Title     field.String
	UpdatedAt field.Time
	CreatedAt field.Time
	Content   field.String
	Like      field.Int64
	Read      field.Int64
	Category  field.String
	Password  field.String
	Tags      articleManyToManyTags

	Authors articleManyToManyAuthors

	Fans articleManyToManyFans

	fieldMap map[string]field.Expr
}

func (a article) Table(newTableName string) *article {
	a.articleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a article) As(alias string) *article {
	a.articleDo.DO = *(a.articleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *article) updateTableName(table string) *article {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Title = field.NewString(table, "title")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.Content = field.NewString(table, "content")
	a.Like = field.NewInt64(table, "like")
	a.Read = field.NewInt64(table, "read")
	a.Category = field.NewString(table, "category")
	a.Password = field.NewString(table, "password")

	a.fillFieldMap()

	return a
}

func (a *article) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *article) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 12)
	a.fieldMap["id"] = a.ID
	a.fieldMap["title"] = a.Title
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["content"] = a.Content
	a.fieldMap["like"] = a.Like
	a.fieldMap["read"] = a.Read
	a.fieldMap["category"] = a.Category
	a.fieldMap["password"] = a.Password

}

func (a article) clone(db *gorm.DB) article {
	a.articleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a article) replaceDB(db *gorm.DB) article {
	a.articleDo.ReplaceDB(db)
	return a
}

type articleManyToManyTags struct {
	db *gorm.DB

	field.RelationField

	Articles struct {
		field.RelationField
	}
}

func (a articleManyToManyTags) Where(conds ...field.Expr) *articleManyToManyTags {
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

func (a articleManyToManyTags) WithContext(ctx context.Context) *articleManyToManyTags {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a articleManyToManyTags) Session(session *gorm.Session) *articleManyToManyTags {
	a.db = a.db.Session(session)
	return &a
}

func (a articleManyToManyTags) Model(m *model.Article) *articleManyToManyTagsTx {
	return &articleManyToManyTagsTx{a.db.Model(m).Association(a.Name())}
}

type articleManyToManyTagsTx struct{ tx *gorm.Association }

func (a articleManyToManyTagsTx) Find() (result []*model.Tag, err error) {
	return result, a.tx.Find(&result)
}

func (a articleManyToManyTagsTx) Append(values ...*model.Tag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a articleManyToManyTagsTx) Replace(values ...*model.Tag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a articleManyToManyTagsTx) Delete(values ...*model.Tag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a articleManyToManyTagsTx) Clear() error {
	return a.tx.Clear()
}

func (a articleManyToManyTagsTx) Count() int64 {
	return a.tx.Count()
}

type articleManyToManyAuthors struct {
	db *gorm.DB

	field.RelationField

	Articles struct {
		field.RelationField
	}
	Likes struct {
		field.RelationField
	}
}

func (a articleManyToManyAuthors) Where(conds ...field.Expr) *articleManyToManyAuthors {
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

func (a articleManyToManyAuthors) WithContext(ctx context.Context) *articleManyToManyAuthors {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a articleManyToManyAuthors) Session(session *gorm.Session) *articleManyToManyAuthors {
	a.db = a.db.Session(session)
	return &a
}

func (a articleManyToManyAuthors) Model(m *model.Article) *articleManyToManyAuthorsTx {
	return &articleManyToManyAuthorsTx{a.db.Model(m).Association(a.Name())}
}

type articleManyToManyAuthorsTx struct{ tx *gorm.Association }

func (a articleManyToManyAuthorsTx) Find() (result []*model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a articleManyToManyAuthorsTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a articleManyToManyAuthorsTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a articleManyToManyAuthorsTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a articleManyToManyAuthorsTx) Clear() error {
	return a.tx.Clear()
}

func (a articleManyToManyAuthorsTx) Count() int64 {
	return a.tx.Count()
}

type articleManyToManyFans struct {
	db *gorm.DB

	field.RelationField

	Articles struct {
		field.RelationField
	}
	Likes struct {
		field.RelationField
	}
}

func (a articleManyToManyFans) Where(conds ...field.Expr) *articleManyToManyFans {
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

func (a articleManyToManyFans) WithContext(ctx context.Context) *articleManyToManyFans {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a articleManyToManyFans) Session(session *gorm.Session) *articleManyToManyFans {
	a.db = a.db.Session(session)
	return &a
}

func (a articleManyToManyFans) Model(m *model.Article) *articleManyToManyFansTx {
	return &articleManyToManyFansTx{a.db.Model(m).Association(a.Name())}
}

type articleManyToManyFansTx struct{ tx *gorm.Association }

func (a articleManyToManyFansTx) Find() (result []*model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a articleManyToManyFansTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a articleManyToManyFansTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a articleManyToManyFansTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a articleManyToManyFansTx) Clear() error {
	return a.tx.Clear()
}

func (a articleManyToManyFansTx) Count() int64 {
	return a.tx.Count()
}

type articleDo struct{ gen.DO }

func (a articleDo) Debug() *articleDo {
	return a.withDO(a.DO.Debug())
}

func (a articleDo) WithContext(ctx context.Context) *articleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleDo) ReadDB() *articleDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleDo) WriteDB() *articleDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleDo) Session(config *gorm.Session) *articleDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleDo) Clauses(conds ...clause.Expression) *articleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleDo) Returning(value interface{}, columns ...string) *articleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleDo) Not(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleDo) Or(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleDo) Select(conds ...field.Expr) *articleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleDo) Where(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleDo) Order(conds ...field.Expr) *articleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleDo) Distinct(cols ...field.Expr) *articleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleDo) Omit(cols ...field.Expr) *articleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleDo) Join(table schema.Tabler, on ...field.Expr) *articleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *articleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleDo) RightJoin(table schema.Tabler, on ...field.Expr) *articleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleDo) Group(cols ...field.Expr) *articleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleDo) Having(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleDo) Limit(limit int) *articleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleDo) Offset(offset int) *articleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *articleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleDo) Unscoped() *articleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleDo) Create(values ...*model.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleDo) CreateInBatches(values []*model.Article, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleDo) Save(values ...*model.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleDo) First() (*model.Article, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Take() (*model.Article, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Last() (*model.Article, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Find() ([]*model.Article, error) {
	result, err := a.DO.Find()
	return result.([]*model.Article), err
}

func (a articleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Article, err error) {
	buf := make([]*model.Article, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleDo) FindInBatches(result *[]*model.Article, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleDo) Attrs(attrs ...field.AssignExpr) *articleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleDo) Assign(attrs ...field.AssignExpr) *articleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleDo) Joins(fields ...field.RelationField) *articleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleDo) Preload(fields ...field.RelationField) *articleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleDo) FirstOrInit() (*model.Article, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) FirstOrCreate() (*model.Article, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) FindByPage(offset int, limit int) (result []*model.Article, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleDo) Delete(models ...*model.Article) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleDo) withDO(do gen.Dao) *articleDo {
	a.DO = *do.(*gen.DO)
	return a
}
