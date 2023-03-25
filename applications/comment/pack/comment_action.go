package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageCommentActionResp(result *commentDomain.DoutokAddCommentResp, e error) (resp *comment.DouyinCommentActionResponse, err error) {
	if e != nil {
		return nil, e
	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		Comment:    result.Comment,
	}, e
}

func PackageCommentActionRespWithErr(errNo *errno.ErrNo) (resp *comment.DouyinCommentActionResponse, err error) {
	return &comment.DouyinCommentActionResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
	}, nil
}
