package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"

	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/applications/comment/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
)

func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	if result := misc.CheckCommentActionArgs(req); !result {
		return pack.PackageCommentActionRespWithErr(&misc.ParamsErr)
	}

	// 判断请求的动作，1为新加评论，2为删除评论
	if req.ActionType == 1 {
		result, err := rpc.CommentDomainRPCClient.AddComment(ctx, &commentDomain.DoutokAddCommentReq{
			VideoId:     req.VideoId,
			UserId:      req.UserId,
			CommentText: req.CommentText,
		})
		return pack.PackageCommentActionResp(result, err)
	} else if req.ActionType == 2 {
		result, err := rpc.CommentDomainRPCClient.RmComment(ctx, &commentDomain.DoutokRmCommentReq{
			VideoId:   req.VideoId,
			UserId:    req.UserId,
			CommentId: req.CommentId,
		})
		return pack.PackageCommentActionResp(result, err)
	} else {
		return pack.PackageCommentActionRespWithErr(&misc.BindingErr)
	}
}
