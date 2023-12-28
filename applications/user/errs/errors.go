package errs

import "github.com/TremblingV5/DouTok/pkg/errno"

var (
	NilErrCode           = -1
	SuccessCode          = 0
	EmptyUserListErrCode = 27001
)

var (
	EmptyUserListErr = errno.New(EmptyUserListErrCode, "Empty user list but downstream errors")
)
