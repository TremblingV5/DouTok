package service

import (
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func CreateFavorite(user_id int64, video_id int64) (*errno.ErrNo, error) {
	op := true

	// 1. 写缓存
	err := WriteFavoriteInCache(user_id, video_id, op)
	if err != nil {
		return &misc.QueryCacheErr, err
	}

	err = UpdateCacheCount(video_id, op)
	if err != nil {
		return &misc.QueryCacheErr, err
	}

	// 2. 写RDB
	err = CreateFavoriteInRDB(user_id, video_id, op)
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
