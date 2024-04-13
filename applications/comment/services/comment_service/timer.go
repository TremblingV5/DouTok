package comment_service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/TremblingV5/DouTok/applications/comment/domain/entity/comment_count"
	"github.com/TremblingV5/DouTok/pkg/dlog"
)

var logger = dlog.InitLog(3)

/*
定时将内存中的局部评论数更新到数据库中，并删除Redis中的评论总数
*/
func (s *Service) UpdateComCountMap() {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		logger.Info("Start iter comment cnt map and update on " + fmt.Sprint(time.Now().Unix()))

		var keyList []string
		s.commentCountCache.Iter(func(key string, v interface{}) {
			keyList = append(keyList, key)

			keyI64, _ := strconv.ParseInt(key, 10, 64)
			err := s.updateCount(context.Background(), keyI64, int64(v.(int)))
			if err != nil {
				dlog.Warn("Write comment count to RDB defeat: " + key + " with count: " + fmt.Sprint(v.(int)))
			}

			err = s.delCount2Cache(key)
			if err != nil {
				dlog.Warn("Delete comment count from third party cache defeat: " + key)
			}
		})

		for _, v := range keyList {
			i64, _ := strconv.ParseInt(v, 10, 64)
			s.commentCountCache.Set(i64, 0)
		}
	}
}

/*
定时将Redis中每个Video的Comment总数更新到内存中的Map
Redis不存在的视频评论数由单独查询时再添加到Redis中
*/
func (s *Service) UpdateComTotalCntMap() {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		logger.Info("Start iter comment total cnt map and update on " + fmt.Sprint(time.Now().Unix()))

		keyList := []string{}

		s.commentTotalCountCache.Iter(func(key string, v interface{}) {
			keyList = append(keyList, key)
		})

		for _, v := range keyList {
			res, err := s.commentTotalCountRedis.Get(context.Background(), v)
			if err != nil {
				continue
			}

			i, _ := strconv.ParseInt(res, 10, 64)
			iKey, _ := strconv.ParseInt(v, 10, 64)
			s.commentTotalCountCache.Set(iKey, i)
		}
	}
}

func (s *Service) updateCount(ctx context.Context, videoId int64, cnt int64) error {
	commentModel, err := s.commentCountRepository.LoadById(ctx, videoId)
	if err != nil {
		return err

	}

	commentCount := comment_count.TransformFromModel(commentModel)

	if err := commentCount.Check(comment_count.IsIdValid()); err != nil {
		return err
	}

	commentCount.UpdateNumber(cnt)

	if err := s.commentCountRepository.Update(ctx, commentCount.ToModel()); err != nil {
		return err
	}

	return nil
}

/*
删除Redis中存储的视频的完整的评论数量
*/
func (s *Service) delCount2Cache(videoId string) error {
	return s.commentTotalCountRedis.Del(context.Background(), videoId)
}
