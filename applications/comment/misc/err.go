package misc

import "github.com/TremblingV5/DouTok/pkg/errno"

var (
	NilErrCode                    = -1
	SuccessCode                   = 0
	ParamsErrCode                 = 26001
	SystemErrCode                 = 26002
	BindingErrCode                = 26003
	ListEmptyErrCode              = 26004
	VideoIdErrCode                = 26005
	QueryCommentCountErrCode      = 26006
	QueryCommentListInHBErrCode   = 26007
	CommentNotBelongToUserErrCode = 26008
	RmDataFromHBErrCode           = 26009
	HBDataNotFoundErrCode         = 26010
)

var (
	NilErr                    = errno.New(NilErrCode, "Don't care")
	Success                   = errno.New(SuccessCode, "Success")
	ParamsErr                 = errno.New(ParamsErrCode, "Invalid parameters")
	SystemErr                 = errno.New(SystemErrCode, "System error")
	BindingErr                = errno.New(BindingErrCode, "Action type invalid")
	ListEmptyErr              = errno.New(ListEmptyErrCode, "List is empty")
	VideoIdErr                = errno.New(VideoIdErrCode, "Video id invalid")
	QueryCommentCountErr      = errno.New(QueryCommentCountErrCode, "Query comment count error")
	QueryCommentListInHBErr   = errno.New(QueryCommentListInHBErrCode, "Query comment list in HBase error")
	CommentNotBelongToUserErr = errno.New(CommentNotBelongToUserErrCode, "This comment not belong to the user")
	RmDataFromHBErr           = errno.New(RmDataFromHBErrCode, "Deleta data in HBase error")
	HBDataNotFoundErr         = errno.New(HBDataNotFoundErrCode, "HBase comment data not found")
)
