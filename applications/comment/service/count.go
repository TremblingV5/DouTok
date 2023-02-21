package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/comment/dal/model"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/go-redis/redis/v8"
	"strconv"
)

/*
	返回一个由video_id组成的列表的评论数统计
*/
func CountComment(video_id ...int64) (map[int64]int64, *errno.ErrNo, error) {
	resMap := make(map[int64]int64)

	// 1. 从内存中查找喜欢数
	findAgain := []int64{}
	for _, v := range video_id {
		cnt, ok, _ := ReadComTotalCount(fmt.Sprint(v))

		if !ok {
			findAgain = append(findAgain, v)
		} else {
			resMap[v] = cnt
		}

	}

	// 2. 从Redis中查找喜欢数
	findAgainAgain := []int64{}
	for _, v := range findAgain {
		cnt, ok, _ := ReadComTotalCountFromCache(fmt.Sprint(v))

		if !ok {
			findAgainAgain = append(findAgainAgain, v)
		} else {
			resMap[v] = cnt
			ComTotalCount.Set(fmt.Sprint(v), cnt)
		}
	}

	// 3. 从MySQL中查找喜欢数
	res, err := DoCommentCnt.Where(
		CommentCnt.Id.In(findAgainAgain...),
	).Find()

	if err != nil {
		return nil, &misc.QueryCommentCountErr, err
	}

	for _, v := range res {
		resMap[v.Id] = v.Number
	}

	// 4. 如果仍然没有查找到该记录，则置0
	for _, v := range video_id {
		if _, ok := resMap[v]; !ok {
			resMap[v] = 0
		}
	}

	return resMap, &misc.Success, nil
}

//func AddCount(video_id int64) error {
//	_, err := DoCommentCnt.Where(
//		CommentCnt.Id.Eq(video_id),
//	).Update(CommentCnt.Number, CommentCnt.Number.Add(1))
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func ReduceCount(video_id int64) error {
//	_, err := DoCommentCnt.Where(
//		CommentCnt.Id.Eq(video_id),
//	).Update(CommentCnt.Number, CommentCnt.Number.Add(-1))
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

/*
	从内存中读取完整的视频评论条数
*/
func ReadComTotalCount(videoId string) (int64, bool, error) {
	data, ok := ComTotalCount.Get(videoId)
	if !ok {
		return 0, false, nil
	}

	return int64(data.(int)), true, nil
}

/*
	从Redis中读取完整的视频评论条数
*/
func ReadComTotalCountFromCache(videoId string) (int64, bool, error) {
	data, err := RedisClients[misc.ComTotalCntCache].Get(context.Background(), videoId)
	if err == redis.Nil {
		return 0, false, nil
	} else if err != nil {
		return 0, false, err
	}

	num, _ := strconv.ParseInt(data, 10, 64)

	return num, true, nil
}

/*
	更新内存中的局部评论计数器
*/
func UpdateCacheComCount(videoId int64, is_add bool) {
	data, ok := ComCount.Get(fmt.Sprint(videoId))

	if ok {
		// 已经存在videoId对应的局部评论数
		if is_add {
			ComCount.Set(fmt.Sprint(videoId), data.(int)+1)
		} else {
			ComCount.Set(fmt.Sprint(videoId), data.(int)-1)
		}
	} else {
		// 尚不存在videoId对应的局部评论数
		if is_add {
			ComCount.Set(fmt.Sprint(videoId), 1)
		} else {
			ComCount.Set(fmt.Sprint(videoId), -1)
		}
	}
}

/*
	更新MySQL中的视频评论条数
*/
func UpdateCount(videoId int64, cnt int64) error {
	data, _ := DoCommentCnt.Where(
		CommentCnt.Id.Eq(videoId),
	).First()

	if data != nil {
		// 已经存在该videoId的计数数据
		_, err := DoCommentCnt.Where(
			CommentCnt.Id.Eq(videoId),
		).Update(CommentCnt.Number, CommentCnt.Number.Add(cnt))

		if err != nil {
			return err
		}
	} else {
		// 新建该videoId的计数数据
		err := DoCommentCnt.Create(
			&model.CommentCount{
				Id:     videoId,
				Number: cnt,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	删除Redis中存储的视频的完整的评论数量
*/
func DelCount2Cache(videoId string) error {
	return RedisClients[misc.ComTotalCntCache].Del(context.Background(), videoId)
}
