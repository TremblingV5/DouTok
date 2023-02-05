package service

import (
	"github.com/TremblingV5/DouTok/applications/favorite/dal/model"
)

/*
点赞关系表结构
type FavRelation struct {
    UserId  uint64
    VideoId uint64
    Status  bool
}
*/

// 记录点赞状态
func SaveFavRelation(user_id int64, video_id int64) error {

	// 2. 写入数据到MySQl
	_, err := SaveFavRelation2DB(
		uint64(user_id), uint64(video_id), bool(true))
	if err != nil {
		return err
	}
	return nil
}

func SaveFavRelation2DB(user_id uint64, video_id uint64, op bool) (uint64, error) {
	newfavRelation := model.FavRelation{
		UserId:  user_id,
		VideoId: video_id,
		Status:  op,
	}

	err := FavRelation.Create(&newfavRelation)

	if err != nil {
		return 0, err
	}

	return newfavRelation.UserId, nil
}
