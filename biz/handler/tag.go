package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"ulog-backend/biz/service"
)

func GetTagList(ctx context.Context, c *app.RequestContext) {
	tagList, err := service.NewTagService(ctx, c).GetTagList()
	if err != nil {
		c.JSON(http.StatusNotFound, buildServiceErrResp(err))
		return
	}
	c.JSON(http.StatusOK, tagList)
}

func AddTag(ctx context.Context, c *app.RequestContext) {
	var req struct {
		Name string `json:"name,required" vd:"$>=0"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, buildParamErrResp(err))
		return
	}
	tag, err := service.NewTagService(ctx, c).AddTag(req.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, buildServiceErrResp(err))
		return
	}
	c.JSON(http.StatusOK, tag)
}
