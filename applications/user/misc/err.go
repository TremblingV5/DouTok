package misc

import "github.com/TremblingV5/DouTok/pkg/errno"

var (
	NilErrCode           = -1
	SuccessCode          = 0
	EmptyUserListErrCode = 27001
)

var (
	NilErr           = errno.NewErrNo(NilErrCode, "Don't care")
	Success          = errno.NewErrNo(SuccessCode, "Success")
	EmptyUserListErr = errno.NewErrNo(EmptyUserListErrCode, "Empty user list but downstream errors")
)
