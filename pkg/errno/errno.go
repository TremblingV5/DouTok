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
)

type ErrNo struct {
	ErrCode int
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	BadRequest             = NewErrNo(BadRequestErrCode, "Request Failed")
<<<<<<< HEAD
	ErrBind                = NewErrNo(ErrBind, "Error occurred while binding the request body to the struct")
=======
	ErrBind                = NewErrNo(ErrBindErrCode, "Error occurred while binding the request body to the struct")
>>>>>>> 04ce33b0b8f7fc06f7d86e0bf6dcb54b331f85e2
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
