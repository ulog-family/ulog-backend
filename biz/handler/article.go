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
