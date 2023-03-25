package misc

import "github.com/TremblingV5/DouTok/pkg/errno"

var (
	NilErrCode    = -1
	SuccessCode   = 0
	EmptyErrCode  = 31001
	SystemErrCode = 31002
)

var (
	NilErr    = errno.NewErrNo(NilErrCode, "Don't care")
	Success   = errno.NewErrNo(SuccessCode, "Success")
	EmptyErr  = errno.NewErrNo(EmptyErrCode, "size of title or data is zero")
	SystemErr = errno.NewErrNo(SystemErrCode, "System error")
)
