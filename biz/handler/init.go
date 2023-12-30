package handler

import (
	"ulog-backend/pkg/errno"
)

type BaseResp struct {
	StatusCode int
	StatusMsg  string
}

func buildParamErrResp(err error) *BaseResp {
	e := errno.ParamErr.WithMessage(err.Error())
	return baseResp(e)
}

func buildServiceErrResp(err error) *BaseResp {
	e := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(e)
}

func baseResp(err errno.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.Code,
		StatusMsg:  err.Msg,
	}
}
