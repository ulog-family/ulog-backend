package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"ulog-backend/biz/dal"
	"ulog-backend/biz/model"
)

var (
	a = dal.Q.Article
)

type ArticleService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewArticleService(ctx context.Context, c *app.RequestContext) *ArticleService {
	return &ArticleService{
		ctx: ctx,
		c:   c,
	}
}

func (articleService ArticleService) GetArticleInfoById(id int64) (*model.Article, error) {
	article, err := a.Omit(a.Content).Where(a.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (articleService ArticleService) GetArticleById(id int64) (*model.Article, error) {
	article, err := a.Preload(a.Authors, a.Tags, a.Fans).Where(a.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (articleService ArticleService) GetArticleList() ([]*model.Article, error) {
	articleList, err := a.Omit(a.Content).Find()
	if err != nil {
		return nil, err
	}
	return articleList, nil
}
