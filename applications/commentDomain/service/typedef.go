package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/hbModel"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/repository/comment"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/repository/commentcnt"
	"github.com/TremblingV5/DouTok/applications/commentDomain/hb/commentHBaseRepo"
	"github.com/TremblingV5/DouTok/applications/commentDomain/redis/commentTotalCountRedis"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"gorm.io/gorm"

	"github.com/TremblingV5/DouTok/applications/commentDomain/cache/commentCntCache"
	"github.com/TremblingV5/DouTok/applications/commentDomain/cache/commentTotalCountCache"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/query"
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

type CommentDomainUtil struct {
	CommentRepository      commentRepository
	CommentCntRepository   commentCntRepository
	CommentHBaseRepository commentHBaseRepository
	CommentTotalCountRedis CommentTotalCountRedis
	CommentCountCache      CommentCntCache        // key: video id, value: modification if count of comments for a video
	CommentTotalCountCache CommentTotalCountCache // key: video id, value: count of comments for a video
	SnowflakeHandle        *utils.SnowflakeHandle
}

//go:generate mockgen -source=typedef.go -destinDBation=./mocks/service_mock.go -package CommentServiceMocks
type IService interface {
	AddComment(ctx context.Context, videoId, userId, conversationId, lastId, toUserId int64, content string) (*entity.Comment, error)
	CountComments(ctx context.Context, videoId ...int64) (map[int64]int64, error)
	ListComment(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error)
	RemoveComment(ctx context.Context, userId, commentId int64) error
}

func NewCommentDomainUtil(
	db *gorm.DB,
	hb *hbaseHandle.HBaseClient,
	commentTotalCountRedis *commentTotalCountRedis.Client,
	snowflakeNode int64,
) *CommentDomainUtil {
	query.SetDefault(db)
	return &CommentDomainUtil{
		CommentRepository:      comment.NewRepository(db),
		CommentCntRepository:   commentcnt.NewRepository(db),
		CommentHBaseRepository: commentHBaseRepo.NewRepository(hb),
		CommentTotalCountRedis: commentTotalCountRedis,
		CommentCountCache:      commentCntCache.NewCache(),
		CommentTotalCountCache: commentTotalCountCache.NewCache(),
		SnowflakeHandle:        utils.NewSnowflakeHandle(snowflakeNode),
	}
}
