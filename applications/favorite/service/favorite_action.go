package service

import (
	"github.com/TremblingV5/DouTok/applications/favorite/dal/model"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"gorm.io/gorm"
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
// 对于取消点赞的，不删对应行信息，而是追加一条新纪录，查询的时候以最新一行信息为准
func SaveFavRelation(req *favorite.DouyinFavoriteActionRequest) error {

	// 2. 写入数据到MySQl
	var op bool
	if req.ActionType == 1 {
		op = true
	}
	if req.ActionType == 2 {
		op = false
	}
	_, err := SaveFavRelation2DB(
		uint64(req.UserId), uint64(req.VideoId), bool(op))
	if err != nil {
		return err
	}
	return nil
}

// 保存点赞状态信息到FavRelation表中,并更改VideoCount表点赞数量
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

	var res *gorm.DB
	if op == true {
		res = DB.Model(VideoCount).Update("fav_count", gorm.Expr("fav_count + ?", 1))
	} else {
		res = DB.Model(VideoCount).Update("fav_count", gorm.Expr("fav_count - ?", 1))
	}

	if res.Error != nil {
		return newfavRelation.UserId, res.Error
	}

	return newfavRelation.UserId, nil
}
