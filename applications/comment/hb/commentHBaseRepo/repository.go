package commentHBaseRepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/comment/dal/hbModel"

	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	tools "github.com/TremblingV5/DouTok/pkg/misc"
)

var (
	CommentsNotFoundByVideoId = errors.New("comments not found by video id")
)

type Repository struct {
	name   string
	client *hbaseHandle.HBaseClient
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go -package CommentHBRepositoryMocks
type IRepository interface {
	Save(ctx context.Context, commentId, videoId, userId, conversationId, lastId, toUserId int64, content string, timestamp string) error
	List(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error)
	Remove(ctx context.Context, videoId, conversationId int64, timestamp string) error
}

func NewRepository(client *hbaseHandle.HBaseClient) *Repository {
	return &Repository{
		name:   "comment",
		client: client,
	}
}

func (r *Repository) Save(ctx context.Context, commentId, videoId, userId, conversationId, lastId, toUserId int64, content string, timestamp string) error {
	hbData := map[string]map[string][]byte{
		"data": {
			"id":              []byte(fmt.Sprint(commentId)),
			"video_id":        []byte(fmt.Sprint(videoId)),
			"user_id":         []byte(fmt.Sprint(userId)),
			"conversation_id": []byte(fmt.Sprint(conversationId)),
			"last_id":         []byte(fmt.Sprint(lastId)),
			"to_user_id":      []byte(fmt.Sprint(toUserId)),
			"content":         []byte(content),
			"timestamp":       []byte(timestamp),
		},
	}

	if err := r.client.Put(
		r.name, getCommentRowKey(videoId, conversationId, "0", timestamp), hbData,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) List(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error) {
	res, err := r.client.Scan(
		r.name, hbaseHandle.GetFilterByRowKeyPrefix(getCommentQueryPrefix(videoId))...,
	)

	if err != nil {
		return nil, CommentsNotFoundByVideoId
	}

	var commentList []*hbModel.CommentInHB

	for _, v := range res {
		temp := hbModel.CommentInHB{}
		err := tools.Map2Struct4HB(v, &temp)
		if err != nil {
			continue
		}
		commentList = append(commentList, &temp)
	}

	return commentList, nil
}

func (r *Repository) Remove(ctx context.Context, videoId, conversationId int64, timestamp string) error {
	return r.client.RmByRowKey(r.name, getCommentRowKey(videoId, conversationId, "0", timestamp))
}
