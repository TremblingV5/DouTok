package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/dal/hbModel"
	"github.com/TremblingV5/DouTok/applications/comment/dal/model"
	"github.com/TremblingV5/DouTok/applications/comment/dal/repository/comment"
	"github.com/TremblingV5/DouTok/applications/comment/dal/repository/commentcnt"
	"github.com/TremblingV5/DouTok/applications/comment/hb/commentHBaseRepo"
	"github.com/TremblingV5/DouTok/applications/comment/redis/commentTotalCountRedis"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"gorm.io/gorm"

	"github.com/TremblingV5/DouTok/applications/comment/cache/commentCntCache"
	"github.com/TremblingV5/DouTok/applications/comment/cache/commentTotalCountCache"
	"github.com/TremblingV5/DouTok/applications/comment/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

type commentRepository interface {
	Save(ctx context.Context, commentId, videoId, userId, conversationId, lastId, toUserId int64, content string, timestamp string) error
	IsCommentFromUser(ctx context.Context, userId, commentId int64) (*model.Comment, bool)
	Remove(ctx context.Context, commentId int64) error
}

type commentHBaseRepository interface {
	Save(ctx context.Context, commentId, videoId, userId, conversationId, lastId, toUserId int64, content string, timestamp string) error
	List(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error)
	Remove(ctx context.Context, videoId, conversationId int64, timestamp string) error
}

type commentCntRepository interface {
	GetCommentsCount(ctx context.Context, videoIdList ...int64) (map[int64]int64, error)
	UpdateCommentsCount(ctx context.Context, videoIdList ...int64) error
}

type CommentCntCache interface {
	Add(videoId int64, modification int64)
	Get(videoId int64) (int64, bool)
	GetAll() map[int64]int64
	Clear()
}

type CommentTotalCountCache interface {
	Get(videoId int64) (int64, bool)
	Set(videoId, count int64)
	SetBatch(batch map[int64]int64)
	Clear()
}

type CommentTotalCountRedis interface {
	Get(ctx context.Context, videoId int64) (int64, error)
	Delete(ctx context.Context, videoId ...int64) error
}

type commentService struct {
	commentRepository      commentRepository
	commentCntRepository   commentCntRepository
	commentHBaseRepository commentHBaseRepository
	commentTotalCountRedis CommentTotalCountRedis
	commentCountCache      CommentCntCache        // key: video id, value: modification if count of comments for a video
	commentTotalCountCache CommentTotalCountCache // key: video id, value: count of comments for a video
	snowflakeHandle        *utils.SnowflakeHandle
}

//go:generate mockgen -source=typedef.go -destinDBation=./mocks/service_mock.go -package CommentServiceMocks
type IService interface {
	AddComment(ctx context.Context, videoId, userId, conversationId, lastId, toUserId int64, content string) (*entity.Comment, error)
	CountComments(ctx context.Context, videoId ...int64) (map[int64]int64, error)
	ListComment(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error)
	RemoveComment(ctx context.Context, userId, commentId int64) error
}

func NewcommentService(
	db *gorm.DB,
	hb *hbaseHandle.HBaseClient,
	commentTotalCountRedis *commentTotalCountRedis.Client,
	snowflakeNode int64,
) *commentService {
	query.SetDefault(db)
	return &commentService{
		commentRepository:      comment.NewRepository(db),
		commentCntRepository:   commentcnt.NewRepository(db),
		commentHBaseRepository: commentHBaseRepo.NewRepository(hb),
		commentTotalCountRedis: commentTotalCountRedis,
		commentCountCache:      commentCntCache.NewCache(),
		commentTotalCountCache: commentTotalCountCache.NewCache(),
		snowflakeHandle:        utils.NewSnowflakeHandle(snowflakeNode),
	}
}
