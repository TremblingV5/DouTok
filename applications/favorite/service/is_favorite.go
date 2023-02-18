package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
)

func QueryIsFavorite(user_id int64, videoIdList []int64) (map[int64]bool, error) {
	resMap := make(map[int64]bool)

	userIdFavoritedInCache, err := RedisClients[misc.FavCache].HGetAll(context.Background(), fmt.Sprint(user_id))
	if err != nil {
		return nil, err
	}

	findAgain := []int64{}

	for _, v := range videoIdList {
		_, ok := userIdFavoritedInCache[fmt.Sprint(v)]
		if ok {
			if userIdFavoritedInCache[fmt.Sprint(v)] == "1" {
				resMap[v] = true
			} else {
				resMap[v] = false
			}
		} else {
			findAgain = append(findAgain, v)
		}
	}

	res, err := DoFavorite.Where(
		Favorite.UserId.Eq(user_id), Favorite.VideoId.In(findAgain...),
	).Find()

	if err != nil {
		return nil, err
	}

	for _, v := range res {
		if v.Status == 1 {
			resMap[v.VideoId] = true
		} else {
			resMap[v.VideoId] = false
		}
	}

	return resMap, nil
}
