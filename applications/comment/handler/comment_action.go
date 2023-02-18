package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/applications/comment/pack"
	"github.com/TremblingV5/DouTok/applications/comment/service"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	if result := misc.CheckCommentActionArgs(req); !result {
		return pack.PackCommentActionResp(int32(misc.ParamsErr.ErrCode), misc.ParamsErr.ErrMsg, nil, req.UserId)
	}

	// 判断请求的动作，1为新加评论，2为删除评论
	if req.ActionType == 1 {
		result, err := service.AddComment(
			req.VideoId, req.UserId, utils.GetSnowFlakeId().Int64(), 0, 0, req.CommentText,
		)
		if err != nil {
			return pack.PackCommentActionResp(int32(misc.SystemErr.ErrCode), misc.SystemErr.ErrMsg, nil, req.UserId)
		}
		return pack.PackCommentActionResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, result, req.UserId)
	} else if req.ActionType == 2 {
		errNo, err := service.RmComment(req.UserId, req.CommentId)
		if err != nil {
			return pack.PackCommentActionResp(int32(errNo.ErrCode), errNo.ErrMsg, nil, req.UserId)
		}
		return pack.PackCommentActionResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, nil, req.UserId)
	} else {
		return pack.PackCommentActionResp(int32(misc.BindingErr.ErrCode), misc.BindingErr.ErrMsg, nil, req.UserId)
	}
}
