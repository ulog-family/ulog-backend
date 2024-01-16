package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"ulog-backend/biz/service"
)

func GetArticleInfoById(ctx context.Context, c *app.RequestContext) {
	var req struct {
		Id int64 `path:"id,required" vd:"$>=0"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, buildParamErrResp(err))
		return
	}
	article, err := service.NewArticleService(ctx, c).GetArticleInfoById(req.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, buildServiceErrResp(err))
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetArticleById(ctx context.Context, c *app.RequestContext) {
	var req struct {
		Id int64 `path:"id,required" vd:"$>=0"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, buildParamErrResp(err))
		return
	}
	article, err := service.NewArticleService(ctx, c).GetArticleById(req.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, buildServiceErrResp(err))
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetArticleList(ctx context.Context, c *app.RequestContext) {
	articleList, err := service.NewArticleService(ctx, c).GetArticleList()
	if err != nil {
		c.JSON(http.StatusNotFound, buildServiceErrResp(err))
		return
	}
	c.JSON(http.StatusOK, articleList)
}

func GetArticleCategoryList(ctx context.Context, c *app.RequestContext) {
	categoryList, err := service.NewArticleService(ctx, c).GetCategoryList()
	if err != nil {
		c.JSON(http.StatusNotFound, buildServiceErrResp(err))
		return
	}
	c.JSON(http.StatusOK, categoryList)
}
