package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	BadRequestErrCode          = 10005
	ErrBindErrCode             = 10006
	InternalErrCode            = 10007
	RedisSetErrorCode          = 10008
	RedisGetErrorCode          = 10009
)

type ErrNo struct {
	ErrCode  int
	ErrMsg   string
	NameCode int
	NodeCode int
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func New(code int, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = New(SuccessCode, "Success")
	ServiceErr             = New(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = New(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = New(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = New(AuthorizationFailedErrCode, "Authorization failed")
	BadRequest             = New(BadRequestErrCode, "Request Failed")
	ErrBind                = New(ErrBindErrCode, "Error occurred while binding the request body to the struct")
	InternalErr            = New(InternalErrCode, "Internal server error")
	RedisSetErr            = New(RedisSetErrorCode, "Set data to redis error")
	RedisGetErr            = New(RedisGetErrorCode, "Get data from redis error")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
