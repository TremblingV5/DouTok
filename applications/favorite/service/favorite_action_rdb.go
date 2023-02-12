package service

import (
	"github.com/TremblingV5/DouTok/applications/favorite/dal/model"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/go-redis/redis/v8"
)

func CreateFavoriteInRDB(user_id int64, video_id int64, is_fav bool) error {
	var op int
	if is_fav {
		op = 1
	} else {
		op = 2
	}

	res, err := DoFavorite.Where(
		Favorite.UserId.Eq(user_id), Favorite.VideoId.Eq(video_id),
	).Find()

	if err != nil && err != redis.Nil {
		return err
	}

	if len(res) > 0 {
		// 已经存在用户与视频之间的关系记录
		if op == res[0].Status {
			return nil
		}

		_, err := DoFavorite.Where(
			Favorite.UserId.Eq(user_id), Favorite.VideoId.Eq(video_id),
		).Update(
			Favorite.Status, op,
		)

		if err != nil {
			return err
		}

		if is_fav {
			AddCount(video_id)
		} else {
			ReduceCount(video_id)
		}

		return nil
	} else {
		// 尚未存在用户与视频之间的关系记录
		id := utils.GetSnowFlakeId()
		err := DoFavorite.Create(
			&model.Favorite{
				ID:      id.Int64(),
				UserId:  user_id,
				VideoId: video_id,
				Status:  op,
			},
		)

		if err != nil {
			return err
		}

		if is_fav {
			AddCount(video_id)
		}

		return nil
	}
}
