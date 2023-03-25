package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	tools "github.com/TremblingV5/DouTok/pkg/misc"
)

type ListCommentService struct {
	ctx context.Context
}

func NewListCommentService(ctx context.Context) *ListCommentService {
	return &ListCommentService{ctx: ctx}
}

/*
	根据video_id列出评论列表
*/
func (s *ListCommentService) ListComment(videoId int64) ([]*model.CommentInHB, errno.ErrNo, error) {
	res, err := HBClient.Scan(
		"comment",
		hbaseHandle.GetFilterByRowKeyPrefix(misc.GetCommentQueryPrefix(videoId))...,
	)

	if err != nil {
		return nil, misc.QueryCommentListInHBErr, err
	}

	comment_list := []*model.CommentInHB{}

	for _, v := range res {
		temp := model.CommentInHB{}
		err := tools.Map2Struct4HB(v, &temp)
		if err != nil {
			continue
		}
		comment_list = append(comment_list, &temp)
	}

	return comment_list, misc.Success, nil
}
