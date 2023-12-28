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
	NilErr                 = errno.New(NilErrCode, "Don't care")
	Success                = errno.New(SuccessCode, "Success")
	ParamsErr              = errno.New(ParamsErrCode, "Invalid parameters")
	SystemErr              = errno.New(SystemErrCode, "System error")
	EmptyVideoIdListErr    = errno.New(EmptyVideoIdListErrCode, "Empty video id list")
	BindingInvalidErr      = errno.New(BindingInvalidErrCode, "Action type is invalid")
	QueryCacheErr          = errno.New(QueryCacheErrCode, "Query cache defeat")
	CreateFavoriteInRDBErr = errno.New(CreateFavoriteInRDBErrCode, "Create new favorite relation in rdb defeat")
	WriteRDBErr            = errno.New(WriteRDBErrCode, "Write rdb defeat")
)
