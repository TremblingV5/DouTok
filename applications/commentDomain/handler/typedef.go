package handler

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/hbModel"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
)

var (
	Success         = "success"
	ParametersError = errors.New("parameters error")
)

type CommentDomainHandler struct {
}

//go:generate mockgen -source=typedef.go -destination=./mocks/service_mock.go -package HandlerServiceMocks
type CommentDomainService interface {
	AddComment(ctx context.Context, videoId, userId, conversationId, lastId, toUserId int64, content string) (*entity.Comment, error)
	CountComments(ctx context.Context, videoId ...int64) (map[int64]int64, error)
	ListComment(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error)
	RemoveComment(ctx context.Context, userId, commentId int64) error
}
