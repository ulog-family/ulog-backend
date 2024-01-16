package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"ulog-backend/biz/dal"
	"ulog-backend/biz/model"
)

var (
	t = dal.Q.Tag
)

type TagService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewTagService(ctx context.Context, c *app.RequestContext) *TagService {
	return &TagService{
		ctx: ctx,
		c:   c,
	}
}

func (tagService TagService) GetTagList() ([]*model.Tag, error) {
	tagList, err := t.Find()
	if err != nil {
		return nil, err
	}
	return tagList, nil
}

func (tagService TagService) AddTag(name string) (*model.Tag, error) {
	newTag := &model.Tag{
		Name: name,
	}
	err := t.Create(newTag)
	if err != nil {
		return nil, err
	}
	return newTag, nil
}
