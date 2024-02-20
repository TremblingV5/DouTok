package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/errs"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

type Handler struct {
	clients *rpc.Clients
}

func New(clients *rpc.Clients) *Handler {
	return &Handler{clients: clients}
}

func (h *Handler) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	if req.CommentText != "" {
		return &comment.DouyinCommentActionResponse{
			StatusCode: errs.ParamsErr.Code(),
			StatusMsg:  errs.ParamsErr.Message(),
		}, nil
	}

	// add comments
	if req.ActionType == 1 {
		result, err := h.clients.Comment.Client.AddComment(ctx, &commentDomain.DoutokAddCommentReq{
			VideoId:     req.VideoId,
			UserId:      req.UserId,
			CommentText: req.CommentText,
		})
		return &comment.DouyinCommentActionResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
		}, err
	}

	// remove comments
	if req.ActionType == 2 {
		result, err := h.clients.Comment.Client.RmComment(ctx, &commentDomain.DoutokRmCommentReq{
			VideoId:   req.VideoId,
			UserId:    req.UserId,
			CommentId: req.CommentId,
		})
		return &comment.DouyinCommentActionResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
		}, err
	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: errs.BindingErr.Code(),
		StatusMsg:  errs.BindingErr.Message(),
	}, nil
}

func (h *Handler) CommentCount(ctx context.Context, req *comment.DouyinCommentCountRequest) (resp *comment.DouyinCommentCountResponse, err error) {
	result, err := h.clients.Comment.Client.CountComment(ctx, &commentDomain.DoutokCountCommentReq{
		VideoIdList: req.VideoIdList,
	})
	return &comment.DouyinCommentCountResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		Result:     result.CommentCount,
	}, nil
}

func (h *Handler) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	result, err := h.clients.Comment.Client.ListComment(ctx, &commentDomain.DoutokListCommentReq{
		VideoId: req.VideoId,
	})

	return &comment.DouyinCommentListResponse{
		StatusCode:  result.StatusCode,
		StatusMsg:   result.StatusMsg,
		CommentList: result.CommentList,
	}, err
}
