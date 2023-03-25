package service

import (
	"context"
	"fmt"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
)

/*
	参数 is_fav 用于描述要写入缓存的关系是怎样的，true表示建立喜欢关系，false表示删除喜欢关系
*/
func WriteFavoriteInCache(user_id int64, video_id int64, is_fav bool) error {
	var op string
	if is_fav {
		op = "1"
	} else {
		op = "2"
	}
	return RedisClients[misc.FavCache].HSet(
		context.Background(), fmt.Sprint(user_id), fmt.Sprint(video_id), op,
	)
}
