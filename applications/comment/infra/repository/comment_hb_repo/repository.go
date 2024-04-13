package comment_hb_repo

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/infra/misc"
	"github.com/TremblingV5/DouTok/applications/comment/infra/model"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	tools "github.com/TremblingV5/DouTok/pkg/misc"
	"github.com/tsuna/gohbase/hrpc"
)

type Repository interface {
	Save(ctx context.Context, comment *model.CommentInHB) error
	ScanByPrefix(ctx context.Context, videoId int64) ([]*model.CommentInHB, error)
	Delete(ctx context.Context, rowKey string) error
}

type PersistRepository struct {
	hbClient *hbaseHandle.HBaseClient
}

func New(hbClient *hbaseHandle.HBaseClient) *PersistRepository {
	return &PersistRepository{
		hbClient: hbClient,
	}
}

func (r *PersistRepository) Save(ctx context.Context, comment *model.CommentInHB) error {
	hbData := map[string]map[string][]byte{
		"data": {
			"id":              comment.Id,
			"video_id":        comment.VideoId,
			"user_id":         comment.UserId,
			"conversation_id": comment.ConversationId,
			"last_id":         comment.LastId,
			"to_user_id":      comment.ToUserId,
			"content":         comment.Content,
			"timestamp":       comment.Timestamp,
		},
	}

	if err := r.hbClient.Put(
		"comment", misc.GetCommentRowKey(
			comment.GetVideoId(), "0",
			comment.GetConversationId(),
			comment.GetTimestamp(),
		), hbData,
	); err != nil {
		return err
	}

	return nil
}

func (r *PersistRepository) ScanByPrefix(ctx context.Context, videoId int64) ([]*model.CommentInHB, error) {
	result, err := r.hbClient.Scan(
		"comment",
		hbaseHandle.GetFilterByRowKeyPrefix(misc.GetCommentQueryPrefix(videoId))...,
	)
	if err != nil {
		return nil, err
	}

	var commentList []*model.CommentInHB

	for _, v := range result {
		temp := model.CommentInHB{}
		err := tools.Map2Struct4HB(v, &temp)
		if err != nil {
			continue
		}
		commentList = append(commentList, &temp)
	}

	return commentList, nil
}

func (r *PersistRepository) Delete(ctx context.Context, rowKey string) error {
	rmReq, err := hrpc.NewDelStr(
		context.Background(),
		"comment", rowKey, make(map[string]map[string][]byte),
	)

	if err != nil {
		return nil
	}

	if _, err := r.hbClient.Client.Delete(rmReq); err != nil {
		return err
	}

	return nil
}
