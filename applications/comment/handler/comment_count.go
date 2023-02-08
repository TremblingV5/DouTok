package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/applications/comment/pack"
	"github.com/TremblingV5/DouTok/applications/comment/service"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
)

func (s *CommentServiceImpl) CommentCount(ctx context.Context, req *comment.DouyinCommentCountRequest) (resp *comment.DouyinCommentCountResponse, err error) {
	if len(req.VideoIdList) == 0 {
		return pack.PackCommentCountResp(int32(misc.ListEmptyErr.ErrCode), misc.ListEmptyErr.ErrMsg, nil)
	}

	res, errNo, err := service.CountComment(req.VideoIdList...)
	if err != nil {
		return pack.PackCommentCountResp(int32(errNo.ErrCode), errNo.ErrMsg, nil)
	}

	return pack.PackCommentCountResp(int32(errNo.ErrCode), errNo.ErrMsg, res)
}
