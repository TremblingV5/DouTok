package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ErrUnknownCode             = 100001
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	BadRequestErrCode          = 10005
	ErrBindErrCode             = 10006
	InternalErrCode            = 10007
	SignatureInvalid           = 401
	TokenInvalid               = 402
	UserAlreadyExist           = 400
	UserNotFound               = 404
	PasswordIncorrect          = 405
	InvalidHash                = 406
	IncompatibleVersion        = 407
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
	ErrUnknown             = NewErrNo(ErrUnknownCode, "Internal server error")
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	BadRequest             = NewErrNo(BadRequestErrCode, "Request Failed")
	ErrBind                = NewErrNo(ErrBindErrCode, "Error occurred while binding the request body to the struct")
	InternalErr            = NewErrNo(InternalErrCode, "Internal server error")
	ErrSignatureInvalid    = NewErrNo(SignatureInvalid, "Signature is invalid")
	ErrTokenInvalid        = NewErrNo(TokenInvalid, "Token invalid")
	ErrUserAlreadyExist    = NewErrNo(UserAlreadyExist, "User already exist")
	ErrUserNotFound        = NewErrNo(UserNotFound, "User not found")
	ErrPasswordIncorrect   = NewErrNo(PasswordIncorrect, "Password was incorrect")
	ErrInvalidHash         = NewErrNo(InvalidHash, "Encoded hash is not in the correct format")
	ErrIncompatibleVersion = NewErrNo(IncompatibleVersion, "Encoded hash is not in the correct format")
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
