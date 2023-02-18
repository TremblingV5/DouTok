package service

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func ActionFavorite(user_id int64, video_id int64, op bool) (*errno.ErrNo, error) {
	// 1. 写缓存
	err := WriteFavoriteInCache(user_id, video_id, op)
	if err != nil {
		return &misc.QueryCacheErr, err
	}

	// err = UpdateCacheCount(video_id, op)
	err = UpdateCacheFavCount(video_id, op)
	if err != nil {
		return &misc.QueryCacheErr, err
	}

	// 2. 写RDB
	// err = CreateFavoriteInRDB(user_id, video_id, op)
	// if err != nil {
	// 	return &misc.WriteRDBErr, err
	// }

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

func RemoveFavorite(user_id int64, video_id int64) (*errno.ErrNo, error) {
	op := false

	// 1. 改缓存
	err := WriteFavoriteInCache(user_id, video_id, op)
	if err != nil {
		return &misc.QueryCacheErr, err
	}

	err = UpdateCacheCount(video_id, op)
	if err != nil {
		return &misc.QueryCacheErr, err
	}

	// 2. 改RDB
	err = CreateFavoriteInRDB(user_id, video_id, op)
	if err != nil {
		return &misc.WriteRDBErr, err
	}

	return &misc.Success, nil
}
