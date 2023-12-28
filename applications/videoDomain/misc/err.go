package misc

import "github.com/TremblingV5/DouTok/pkg/errno"

var (
	NilErrCode    = -1
	SuccessCode   = 0
	EmptyErrCode  = 31001
	SystemErrCode = 31002
)

var (
	NilErr    = errno.New(NilErrCode, "Don't care")
	Success   = errno.New(SuccessCode, "Success")
	EmptyErr  = errno.New(EmptyErrCode, "size of title or data is zero")
	SystemErr = errno.New(SystemErrCode, "System error")
)
