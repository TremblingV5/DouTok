package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/applications/comment/pack"
	"github.com/TremblingV5/DouTok/applications/comment/service"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
)

func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	if !misc.CheckCommentListArgs(req) {
		return pack.PackCommentListResp(int32(misc.VideoIdErr.ErrCode), misc.VideoIdErr.ErrMsg, nil)
	}

	res, errNo, err := service.ListComment(req.VideoId)
	if err != nil {
		return pack.PackCommentListResp(int32(errNo.ErrCode), errNo.ErrMsg, nil)
	}

	return pack.PackCommentListResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, res)
}
