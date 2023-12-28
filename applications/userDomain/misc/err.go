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
	NilErr             = errno.New(NilErrCode, "Don't care")
	Success            = errno.New(SuccessCode, "Success")
	UserNameErr        = errno.New(UserNameErrCode, "Username error")
	PasswordErr        = errno.New(PasswordErrCode, "Password error")
	EmptyErr           = errno.New(EmptyErrCode, "Username or password is empty")
	UserNameExistedErr = errno.New(UserNameExistedErrCode, "Username existed")
	SearchErr          = errno.New(SearchErrCode, "Search defeat")
	SystemErr          = errno.New(SystemErrCode, "System error")
)
