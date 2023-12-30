package errno

import "fmt"

type ResponseErrType string

type ErrNo struct {
	Code         int             // 业务错误码
	Msg          string          // 错误信息
	ResponseType ResponseErrType // 响应类型
	Cause        error           // 原始错误
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.Code, e.Msg)
}

func NewErrNo(code int, msg string) ErrNo {
	return ErrNo{
		Code: code,
		Msg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.Msg = msg
	return e
}

const (
	ServiceErrCode = iota + 10000
	ParamErrCode
	AlreadyExistCode
)

const (
	ServiceErrMsg   = "Service error"
	ParamErrMsg     = "Param error"
	AlreadyExistMsg = "Already exist"
)

var (
	ServiceErr      = NewErrNo(ServiceErrCode, ServiceErrMsg)
	ParamErr        = NewErrNo(ParamErrCode, ParamErrMsg)
	AlreadyExistErr = NewErrNo(AlreadyExistCode, AlreadyExistMsg)
)
