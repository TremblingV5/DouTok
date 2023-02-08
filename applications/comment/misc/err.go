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
	NilErr                    = errno.NewErrNo(NilErrCode, "Don't care")
	Success                   = errno.NewErrNo(SuccessCode, "Success")
	ParamsErr                 = errno.NewErrNo(ParamsErrCode, "Invalid parameters")
	SystemErr                 = errno.NewErrNo(SystemErrCode, "System error")
	BindingErr                = errno.NewErrNo(BindingErrCode, "Action type invalid")
	ListEmptyErr              = errno.NewErrNo(ListEmptyErrCode, "List is empty")
	VideoIdErr                = errno.NewErrNo(VideoIdErrCode, "Video id invalid")
	QueryCommentCountErr      = errno.NewErrNo(QueryCommentCountErrCode, "Query comment count error")
	QueryCommentListInHBErr   = errno.NewErrNo(QueryCommentListInHBErrCode, "Query comment list in HBase error")
	CommentNotBelongToUserErr = errno.NewErrNo(CommentNotBelongToUserErrCode, "This comment not belong to the user")
	RmDataFromHBErr           = errno.NewErrNo(RmDataFromHBErrCode, "Deleta data in HBase error")
	HBDataNotFoundErr         = errno.NewErrNo(HBDataNotFoundErrCode, "HBase comment data not found")
)
