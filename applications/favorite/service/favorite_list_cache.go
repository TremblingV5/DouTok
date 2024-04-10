package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
)

func QueryFavListInCache(user_id int64) ([]int64, error) {
	res, err := RedisClients[misc.FavCache].HGetAll(context.Background(), fmt.Sprint(user_id))

	if err != nil {
		return nil, err
	}

	result := []int64{}
	for k, v := range res {
		k_i64, _ := strconv.ParseInt(k, 10, 64)
		if v == "1" {
			result = append(result, k_i64)
		}
	}

	return result, nil
}

func WriteFavListInCache(user_id int64, list []int64) error {
	var op []string

	for _, v := range list {
		op = append(op, fmt.Sprint(v))
		op = append(op, "1")
	}

	return RedisClients[misc.FavCache].HSetMore(context.Background(), fmt.Sprint(user_id), op...)
}
