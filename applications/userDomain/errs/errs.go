package errs

import "github.com/TremblingV5/DouTok/pkg/response"

const (
	userNameError       = "username error"
	userNameErrorCode   = 101
	passwordError       = "password error"
	passwordErrorCode   = 102
	emptyError          = "username or password is empty"
	emptyErrorCode      = 103
	userNameExisted     = "username existed"
	userNameExistedCode = 104
	systemError         = "system error"
	systemErrorCode     = 105
)

var (
	Success            *response.Response
	UserNameErr        *response.Response
	PasswordErr        *response.Response
	EmptyErr           *response.Response
	UserNameExistedErr *response.Response
	SystemErr          *response.Response
)

func Init(config response.Config) {
	Success = response.Success(config)

	UserNameErr = Success.Copy().Update(response.Code(userNameErrorCode), response.Message(userNameError))
	PasswordErr = Success.Copy().Update(response.Code(passwordErrorCode), response.Message(passwordError))
	EmptyErr = Success.Copy().Update(response.Code(emptyErrorCode), response.Message(emptyError))
	UserNameExistedErr = Success.Copy().Update(response.Code(userNameExistedCode), response.Message(userNameExisted))
	SystemErr = Success.Copy().Update(response.Code(systemErrorCode), response.Message(systemError))
}
