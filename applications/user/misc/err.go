package misc

import "github.com/TremblingV5/DouTok/pkg/errno"

var (
	NilErrCode             = -1
	SuccessCode            = 0
	UserNameErrCode        = 27001
	PasswordErrCode        = 27002
	EmptyErrCode           = 27003
	UserNameExistedErrCode = 27004
	SearchErrCode          = 27005
	SystemErrCode          = 27999
)

var (
	NilErr             = errno.NewErrNo(NilErrCode, "Don't care")
	Success            = errno.NewErrNo(SuccessCode, "Success")
	UserNameErr        = errno.NewErrNo(UserNameErrCode, "Username error")
	PasswordErr        = errno.NewErrNo(PasswordErrCode, "Password error")
	EmptyErr           = errno.NewErrNo(EmptyErrCode, "Username or password is empty")
	UserNameExistedErr = errno.NewErrNo(UserNameExistedErrCode, "Username existed")
	SearchErr          = errno.NewErrNo(SearchErrCode, "Search defeat")
	SystemErr          = errno.NewErrNo(SystemErrCode, "System error")
)
