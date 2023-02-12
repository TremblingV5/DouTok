package misc

import "github.com/TremblingV5/DouTok/pkg/errno"

var (
	NilErrCode                 = -1
	SuccessCode                = 0
	ParamsErrCode              = 27001
	SystemErrCode              = 27002
	EmptyVideoIdListErrCode    = 27003
	BindingInvalidErrCode      = 27004
	QueryCacheErrCode          = 27005
	CreateFavoriteInRDBErrCode = 27006
	WriteRDBErrCode            = 27007
)

var (
	NilErr                 = errno.NewErrNo(NilErrCode, "Don't care")
	Success                = errno.NewErrNo(SuccessCode, "Success")
	ParamsErr              = errno.NewErrNo(ParamsErrCode, "Invalid parameters")
	SystemErr              = errno.NewErrNo(SystemErrCode, "System error")
	EmptyVideoIdListErr    = errno.NewErrNo(EmptyVideoIdListErrCode, "Empty video id list")
	BindingInvalidErr      = errno.NewErrNo(BindingInvalidErrCode, "Action type is invalid")
	QueryCacheErr          = errno.NewErrNo(QueryCacheErrCode, "Query cache defeat")
	CreateFavoriteInRDBErr = errno.NewErrNo(CreateFavoriteInRDBErrCode, "Create new favorite relation in rdb defeat")
	WriteRDBErr            = errno.NewErrNo(WriteRDBErrCode, "Write rdb defeat")
)
