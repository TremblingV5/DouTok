package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/hbModel"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/go-redis/redis/v8"
	"golang.org/x/exp/maps"
)

var (
	CommentNotBelongToUser = errors.New("comment not belong to user")
	DomainUtil             *CommentDomainUtil
)

func (s *CommentDomainUtil) AddComment(
	ctx context.Context,
	videoId, userId, conversationId, lastId, toUserId int64,
	content string,
) (*entity.Comment, error) {
	timestamp := fmt.Sprint(time.Now().Unix())
	id := s.SnowflakeHandle.GetId()

	result := entity.Comment{
		Id:         int64(id),
		Content:    content,
		CreateDate: "刚刚",
	}

	if err := s.CommentRepository.Save(ctx, id.Int64(), videoId, userId, conversationId, lastId, toUserId, content, timestamp); err != nil {
		return nil, err
	}

	if err := s.CommentHBaseRepository.Save(ctx, id.Int64(), videoId, userId, conversationId, lastId, toUserId, content, timestamp); err != nil {
		return nil, err
	}

	s.CommentCountCache.Add(videoId, 1)

	return &result, nil
}

func (s *CommentDomainUtil) CountComments(ctx context.Context, videoId ...int64) (map[int64]int64, error) {
	resultMap := make(map[int64]int64)

	// Query comments count from cache
	queryAgain := []int64{}
	for _, v := range videoId {
		cnt, ok := s.CommentCountCache.Get(v)
		if !ok {
			queryAgain = append(queryAgain, v)
		} else {
			resultMap[v] = cnt
		}
	}

	// Query comments count from redis
	query3times := []int64{}
	for _, v := range queryAgain {
		cnt, err := s.CommentTotalCountRedis.Get(ctx, v)
		if err != nil && err != redis.Nil {
			return nil, err
		} else if err == redis.Nil {
			query3times = append(query3times, v)
		} else {
			resultMap[v] = cnt
			s.CommentCountCache.Add(v, cnt)
		}
	}

	// Query comments count from mysql
	results, err := s.CommentCntRepository.GetCommentsCount(ctx, query3times...)
	if err != nil {
		return resultMap, err
	}

	maps.Copy(resultMap, results)
	return resultMap, nil
}

func (s *CommentDomainUtil) ListComment(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error) {
	return s.CommentHBaseRepository.List(ctx, videoId)
}

func (s *CommentDomainUtil) RemoveComment(ctx context.Context, userId, commentId int64) error {
	comment, ok := s.CommentRepository.IsCommentFromUser(ctx, userId, commentId)
	if !ok {
		return CommentNotBelongToUser
	}

	if err := s.CommentRepository.Remove(ctx, commentId); err != nil {
		return err
	}

	if err := s.CommentHBaseRepository.Remove(ctx, comment.VideoId, comment.ConversationId, comment.Timestamp); err != nil {
		return err
	}

	s.CommentCountCache.Add(commentId, -1)

	return nil
}
