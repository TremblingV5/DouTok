package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"

	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func ActionFavorite(user_id int64, video_id int64, op bool) (*errno.ErrNo, error) {
	// 1. 在Redis中查找是否已经存在点赞关系的缓存
	_, err := RedisClients[misc.FavCache].HGetAll(context.Background(), fmt.Sprint(user_id))

	if err == redis.Nil {
		// 如果不存在该用户的点赞关系缓存，则从数据库中读出并加载到缓存中
		res, err1 := QueryFavListInRDB(user_id)

		if err1 != nil {
			return &misc.SystemErr, err1
		}

		for _, v := range res {
			if err := WriteFavoriteInCache(user_id, v, true); err != nil {
				continue
			}
		}
	} else if err != nil {
		return &misc.SystemErr, err
	}

	existed, err := RedisClients[misc.FavCache].HGet(context.Background(), fmt.Sprint(user_id), fmt.Sprint(video_id))
	if err != nil && err != redis.Nil {
		return &misc.SystemErr, err
	}

	// 2. 写缓存
	err = WriteFavoriteInCache(user_id, video_id, op)
	if err != nil {
		return &misc.QueryCacheErr, err
	}

	// 3. 更新内存中的计数Map
	if (existed == "1" && op == false) || (existed == "2" && op == true) {
		err = UpdateCacheFavCount(video_id, op)
		if err != nil {
			return &misc.QueryCacheErr, err
		}
	}

	// 4. 通过Kafka延迟落库
	msg := FavReqInKafka{
		UserId:  user_id,
		VideoId: video_id,
		Op:      op,
	}
	json_msg, _ := json.Marshal(msg)

	_, _, err = FavCountKafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: misc.FavCountTopicName,
		Key:   sarama.StringEncoder(json_msg),
		Value: sarama.StringEncoder(json_msg),
	})
	if err != nil {
		return &misc.WriteRDBErr, err
	}

	return &misc.Success, nil
}
