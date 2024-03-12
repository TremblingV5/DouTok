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
)

func (s *CommentDomainService) AddComment(
	ctx context.Context,
	videoId, userId, conversationId, lastId, toUserId int64,
	content string,
) (*entity.Comment, error) {
	timestamp := fmt.Sprint(time.Now().Unix())
	id := s.snowflakeHandle.GetId()

	result := entity.Comment{
		Id:         int64(id),
		Content:    content,
		CreateDate: "刚刚",
	}

	if err := s.commentRepository.Save(ctx, id.Int64(), videoId, userId, conversationId, lastId, toUserId, content, timestamp); err != nil {
		return nil, err
	}

	if err := s.commentHBaseRepository.Save(ctx, id.Int64(), videoId, userId, conversationId, lastId, toUserId, content, timestamp); err != nil {
		return nil, err
	}

	s.commentCountCache.Add(videoId, 1)

	return &result, nil
}

func (s *CommentDomainService) CountComments(ctx context.Context, videoId ...int64) (map[int64]int64, error) {
	resultMap := make(map[int64]int64)

	// Query comments count from cache
	queryAgain := []int64{}
	for _, v := range videoId {
		cnt, ok := s.commentCountCache.Get(v)
		if !ok {
			queryAgain = append(queryAgain, v)
		} else {
			resultMap[v] = cnt
		}
	}

	// Query comments count from redis
	query3times := []int64{}
	for _, v := range queryAgain {
		cnt, err := s.commentTotalCountRedis.Get(ctx, v)
		if err != nil && err != redis.Nil {
			return nil, err
		} else if err == redis.Nil {
			query3times = append(query3times, v)
		} else {
			resultMap[v] = cnt
			s.commentCountCache.Add(v, cnt)
		}
	}

	// Query comments count from mysql
	results, err := s.commentCntRepository.GetCommentsCount(ctx, query3times...)
	if err != nil {
		return resultMap, err
	}

	maps.Copy(resultMap, results)
	return resultMap, nil
}

func (s *CommentDomainService) ListComment(ctx context.Context, videoId int64) ([]*hbModel.CommentInHB, error) {
	return s.commentHBaseRepository.List(ctx, videoId)
}

func (s *CommentDomainService) RemoveComment(ctx context.Context, userId, commentId int64) error {
	comment, ok := s.commentRepository.IsCommentFromUser(ctx, userId, commentId)
	if !ok {
		return CommentNotBelongToUser
	}

	if err := s.commentRepository.Remove(ctx, commentId); err != nil {
		return err
	}

	if err := s.commentHBaseRepository.Remove(ctx, comment.VideoId, comment.ConversationId, comment.Timestamp); err != nil {
		return err
	}

	s.commentCountCache.Add(commentId, -1)

	return nil
}
