package dal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"ulog-backend/config"
)

func init() {
	dbConfig := config.DB
	g := gen.NewGenerator(gen.Config{
		OutPath:       "biz/dal", // output directory, default value is ./query
		ModelPkgPath:  "biz/model",
		Mode:          gen.WithDefaultQuery | gen.WithoutContext,
		FieldNullable: true,
	})
	db, err := gorm.Open(postgres.Open(dbConfig.DSN()), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	//err = db.AutoMigrate(&model.Category{}, &model.Article{}, &model.Tag{}, &model.User{})
	//if err != nil {
	//	panic("迁移数据表失败")
	//}
	g.UseDB(db)
	tagModel := g.GenerateModel("tag",
		gen.FieldRelate(field.Many2Many, "Articles", g.GenerateModel("article"), &field.RelateConfig{
			RelateSlice: true,
			GORMTag:     field.GormTag{}.Set("many2many", "articles_tags"),
			JSONTag:     "articles,omitempty",
		}))
	userModel := g.GenerateModel("user",
		gen.FieldRelate(field.Many2Many, "Articles", g.GenerateModel("article"), &field.RelateConfig{
			RelateSlice: true,
			GORMTag:     field.GormTag{}.Set("many2many", "articles_authors"),
			JSONTag:     "articles,omitempty",
		}),
		gen.FieldRelate(field.Many2Many, "Likes", g.GenerateModel("article"), &field.RelateConfig{
			RelateSlice: true,
			GORMTag:     field.GormTag{}.Set("many2many", "articles_likes"),
			JSONTag:     "likes,omitempty",
		}),
		gen.FieldNewTag("password", field.Tag{"json": "-"}),
		gen.FieldNewTag("email", field.Tag{"json": "email,omitempty"}),
		gen.FieldNewTag("birthday", field.Tag{"json": "birthday,omitempty"}),
		gen.FieldNewTag("articles", field.Tag{"json": "articles,omitempty"}),
		gen.FieldNewTag("created_at", field.Tag{"json": "-"}),
		gen.FieldNewTag("updated_at", field.Tag{"json": "-"}),
	)
	articleModel := g.GenerateModel("article",
		gen.FieldIgnore("password"),
		gen.FieldRelate(field.Many2Many, "Tags", tagModel, &field.RelateConfig{
			RelateSlice: true,
			GORMTag:     field.GormTag{}.Set("many2many", "articles_tags"),
			JSONTag:     "tags,omitempty",
		}),
		gen.FieldRelate(field.Many2Many, "Authors", userModel, &field.RelateConfig{
			RelateSlice: true,
			GORMTag:     field.GormTag{}.Set("many2many", "articles_authors"),
		}),
		gen.FieldRelate(field.Many2Many, "Fans", userModel, &field.RelateConfig{
			RelateSlice: true,
			GORMTag:     field.GormTag{}.Set("many2many", "articles_likes"),
			JSONTag:     "fans,omitempty",
		}),
	)
	g.ApplyBasic(tagModel, articleModel, userModel)
	g.Execute()
	//SetDefault(db)
}
