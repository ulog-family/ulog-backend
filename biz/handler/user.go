package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"ulog-backend/biz/service"
)

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req struct {
		Name string `path:"name,required"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, buildParamErrResp(err))
		return
	}
	user, err := service.NewUserService(ctx, c).GetUserInfoByName(req.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, buildServiceErrResp(err))
		return
	}
	c.JSON(http.StatusOK, user)
}

func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var req struct {
		Method   string `path:"method"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, buildParamErrResp(err))
		return
	}
	switch req.Method {
	case "passkey":
		return
	default:
		user, err := service.NewUserService(ctx, c).PasswordRegisterUser(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func AuthenticateUser(ctx context.Context, c *app.RequestContext) (any, error) {
	var req struct {
		Method   string `path:"method"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, buildParamErrResp(err))
		return nil, err
	}
	switch req.Method {
	case "passkey":
		return nil, nil
	default:
		user, err := service.NewUserService(ctx, c).PasswordAuthenticateUser(req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}
