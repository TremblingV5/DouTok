package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/service"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/pkg/utils"

	"github.com/TremblingV5/DouTok/applications/comment/errs"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
)

type Handler struct {
	comment service.IService
}

func New(
	comment service.IService,
) *Handler {
	return &Handler{
		comment: comment,
	}
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
		result, err := h.comment.AddComment(ctx, req.VideoId, req.UserId, utils.GetSnowFlakeId().Int64(), 0, 0, req.CommentText)
		if err != nil {
			return &comment.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			}, err
		}

		return &comment.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    result,
		}, nil
	}

	// remove comments
	if req.ActionType == 2 {
		err := h.comment.RemoveComment(ctx, req.UserId, req.CommentId)
		if err != nil {
			return &comment.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			}, err
		}

	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: errs.BindingErr.Code(),
		StatusMsg:  errs.BindingErr.Message(),
	}, nil
}

func (h *Handler) CommentCount(ctx context.Context, req *comment.DouyinCommentCountRequest) (resp *comment.DouyinCommentCountResponse, err error) {
	result, err := h.comment.CountComments(ctx, req.VideoIdList...)

	if err != nil {
		return nil, err
	}

	return &comment.DouyinCommentCountResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Result:     result,
	}, nil
}

func (h *Handler) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	result, err := h.comment.ListComment(ctx, req.VideoId)
	if err != nil {
		return &comment.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		}, err
	}

	// TODO: 完成解耦后增加完成的数据聚合
	var commentList []*entity.Comment
	for _, item := range result {
		commentList = append(commentList, &entity.Comment{
			Id:         item.GetId(),
			Content:    item.GetContent(),
			CreateDate: item.GetTimestamp(),
		})
	}

	return &comment.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: commentList,
	}, err
}
