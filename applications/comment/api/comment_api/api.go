package comment_api

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/infra/misc"
	"github.com/TremblingV5/DouTok/applications/comment/infra/rpc"

	"github.com/TremblingV5/DouTok/applications/comment/services/comment_service"

	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

type CommentServiceImpl struct {
	commentService *comment_service.Service
	clients        *rpc.Clients
}

func New(commentService *comment_service.Service, clients *rpc.Clients) *CommentServiceImpl {
	return &CommentServiceImpl{
		commentService: commentService,
		clients:        clients,
	}
}

func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	if result := CheckCommentActionArgs(req); !result {
		return s.packCommentActionResp(int32(misc.ParamsErr.ErrCode), misc.ParamsErr.ErrMsg, nil, req.UserId)
	}

	// 判断请求的动作，1为新加评论，2为删除评论
	if req.ActionType == 1 {
		result, err := s.commentService.AddComment(
			ctx, req.VideoId, req.UserId, utils.GetSnowFlakeId().Int64(), 0, 0, req.CommentText,
		)
		if err != nil {
			return s.packCommentActionResp(int32(misc.SystemErr.ErrCode), misc.SystemErr.ErrMsg, nil, req.UserId)
		}
		return s.packCommentActionResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, result, req.UserId)
	} else if req.ActionType == 2 {
		errNo, err := s.commentService.RmComment(ctx, req.UserId, req.CommentId)
		if err != nil {
			return s.packCommentActionResp(int32(errNo.ErrCode), errNo.ErrMsg, nil, req.UserId)
		}
		return s.packCommentActionResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, nil, req.UserId)
	} else {
		return s.packCommentActionResp(int32(misc.BindingErr.ErrCode), misc.BindingErr.ErrMsg, nil, req.UserId)
	}
}

func (s *CommentServiceImpl) CommentCount(ctx context.Context, req *comment.DouyinCommentCountRequest) (resp *comment.DouyinCommentCountResponse, err error) {
	if len(req.VideoIdList) == 0 {
		return s.packCommentCountResp(int32(misc.ListEmptyErr.ErrCode), misc.ListEmptyErr.ErrMsg, nil)
	}

	res, errNo, err := s.commentService.CountComment(ctx, req.VideoIdList...)
	if err != nil {
		return s.packCommentCountResp(int32(errNo.ErrCode), errNo.ErrMsg, nil)
	}

	return s.packCommentCountResp(int32(errNo.ErrCode), errNo.ErrMsg, res)
}

func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	if !CheckCommentListArgs(req) {
		return s.packCommentListResp(int32(misc.VideoIdErr.ErrCode), misc.VideoIdErr.ErrMsg, nil)
	}

	res, errNo, err := s.commentService.ListComment(ctx, req.VideoId)
	if err != nil {
		return s.packCommentListResp(int32(errNo.ErrCode), errNo.ErrMsg, nil)
	}

	return s.packCommentListResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, res)
}
