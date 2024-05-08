package comment_service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	misc "github.com/TremblingV5/DouTok/applications/comment/infra/misc"
	"github.com/TremblingV5/box/dbtx"

	"github.com/TremblingV5/DouTok/applications/comment/domain/entity/comment"
	"github.com/TremblingV5/DouTok/applications/comment/domain/entity/comment_count"
	"github.com/TremblingV5/DouTok/applications/comment/infra/repository/comment_count_repo"
	"github.com/TremblingV5/DouTok/applications/comment/infra/repository/comment_hb_repo"
	"github.com/TremblingV5/DouTok/applications/comment/infra/repository/comment_repo"
	"github.com/TremblingV5/DouTok/pkg/cache"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/go-redis/redis/v8"

	"github.com/TremblingV5/DouTok/pkg/utils"
)

type Service struct {
	commentCountCache      *cache.CountMapCache
	commentTotalCountCache *cache.CountMapCache
	commentHBRepository    comment_hb_repo.Repository
	commentCountRedis      *redis.Client
	commentTotalCountRedis *redis.Client
	commentRepository      comment_repo.Repository
	commentCountRepository comment_count_repo.Repository
}

func New(
	commentCountCache *cache.CountMapCache,
	commentTotalCountCache *cache.CountMapCache,
	commentHBRepository comment_hb_repo.Repository,
	commentCountRedis *redis.Client,
	commentTotalCountRedis *redis.Client,
) *Service {
	return &Service{
		commentCountCache:      commentCountCache,
		commentTotalCountCache: commentTotalCountCache,
		commentHBRepository:    commentHBRepository,
		commentCountRedis:      commentCountRedis,
		commentTotalCountRedis: commentTotalCountRedis,
		commentRepository:      comment_repo.New(),
		commentCountRepository: comment_count_repo.New(),
	}
}

func (s *Service) AddComment(
	ctx context.Context,
	videoId int64, userId int64, conId int64,
	lastId int64, toUserId int64, content string,
) (c *comment.Entity, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	timestamp := fmt.Sprint(time.Now().Unix())

	id := utils.GetSnowFlakeId().Int64()

	comment, err := comment.New(
		comment.WithId(id),
		comment.WithVideoId(videoId),
		comment.WithUserId(userId),
		comment.WithConversationId(conId),
		comment.WithLastId(lastId),
		comment.WithToUserId(toUserId),
		comment.WithContent(content),
		comment.WithTimestamp(timestamp),
	)

	if err != nil {
		return nil, err
	}

	if err := s.commentRepository.Save(ctx, comment.ToModel()); err != nil {
		return nil, err
	}

	// TODO: 通过消息队列异步写
	if err := s.commentHBRepository.Save(ctx, comment.ToHBModel()); err != nil {
		return nil, err
	}

	s.updateCacheComCount(videoId, true)

	return comment, nil
}

func (s *Service) CountComment(ctx context.Context, videoId ...int64) (result map[int64]int64, e *errno.ErrNo, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	resMap := make(map[int64]int64)

	// 1. 从内存中查找喜欢数
	var findAgain []int64
	for _, v := range videoId {
		cnt, ok := s.readComTotalCount(v)

		if !ok {
			findAgain = append(findAgain, v)
			continue
		}

		resMap[v] = cnt
	}

	// 2. 从Redis中查找喜欢数
	var findAgainAgain []int64
	for _, v := range findAgain {
		cnt, ok, _ := s.readComTotalCountFromCache(fmt.Sprint(v))

		if !ok {
			findAgainAgain = append(findAgainAgain, v)
			continue
		}

		resMap[v] = cnt
		s.commentTotalCountCache.Set(v, cnt)
	}

	// 3. 从MySQL中查找喜欢数
	commentCountModels, err := s.commentCountRepository.LoadByIdList(ctx, findAgainAgain...)
	if err != nil {
		return nil, &misc.QueryCommentCountErr, err
	}

	commentCounts := comment_count.NewListFromModels(commentCountModels)
	for _, v := range commentCounts {
		resMap[v.Id] = v.Number
	}

	// 4. 如果仍然没有查找到该记录，则置0
	for _, v := range videoId {
		if _, ok := resMap[v]; !ok {
			resMap[v] = 0
		}
	}

	return resMap, &misc.Success, nil
}

func (s *Service) readComTotalCount(videoId int64) (int64, bool) {
	return s.commentTotalCountCache.Get(videoId)
}

func (s *Service) readComTotalCountFromCache(videoId string) (int64, bool, error) {
	data, err := s.commentTotalCountRedis.Get(context.Background(), videoId).Result()
	if errors.Is(err, redis.Nil) {
		return 0, false, nil
	} else if err != nil {
		return 0, false, err
	}

	num, _ := strconv.ParseInt(data, 10, 64)

	return num, true, nil
}

func (s *Service) updateCacheComCount(videoId int64, isAdd bool) {
	if isAdd {
		s.commentCountCache.Add(videoId, 1)
		return
	}

	s.commentCountCache.Add(videoId, -1)
}

func (s *Service) ListComment(ctx context.Context, videoId int64) (comment.List, errno.ErrNo, error) {
	commentHBModels, err := s.commentHBRepository.ScanByPrefix(ctx, videoId)
	if err != nil {
		return nil, misc.QueryCommentListInHBErr, err
	}

	commentList := comment.NewListFromHBModels(commentHBModels)
	return commentList, misc.Success, nil
}

func (s *Service) RmComment(ctx context.Context, userId, commentId int64) (e errno.ErrNo, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	commentModel, err := s.commentRepository.LoadById(ctx, commentId)
	if err != nil {
		return misc.SystemErr, err
	}

	comment, err := comment.New(
		comment.WithId(commentModel.Id),
		comment.WithVideoId(commentModel.VideoId),
		comment.WithUserId(commentModel.UserId),
		comment.WithConversationId(commentModel.ConversationId),
		comment.WithLastId(commentModel.LastId),
		comment.WithToUserId(commentModel.ToUserId),
		comment.WithContent(commentModel.Content),
		comment.WithTimestamp(commentModel.Timestamp),
		comment.WithStatus(commentModel.Status),
		comment.WithCreatedAt(commentModel.CreatedAt),
		comment.WithUpdatedAt(commentModel.UpdatedAt),
	)
	if err != nil {
		return misc.SystemErr, err
	}

	_, err = comment.IsBelongUser(userId)
	if err != nil {
		return misc.SystemErr, err
	}

	comment.MarkAsDelete()

	if err := s.commentRepository.Update(ctx, comment.ToModel()); err != nil {
		return misc.SystemErr, err
	}

	if err := s.commentHBRepository.Save(ctx, comment.ToHBModel()); err != nil {
		return misc.SystemErr, err
	}

	s.updateCacheComCount(comment.VideoId, false)

	return misc.Success, nil
}
