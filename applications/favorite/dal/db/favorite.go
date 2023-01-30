package db

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/constants"

	"gorm.io/gorm"
)

// GetFavoriteRelation get favorite video info
// 会用到feed和user里的表结构
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (*Video, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	video := new(Video)
	// if err := DB.WithContext(ctx).First(&video, vid).Error; err != nil {
	// 	return nil, err
	// }

	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&video, vid); err != nil {
		return nil, err
	}
	return video, nil
}

// Favorite new favorite data.
func Favorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作
		//1. 新增点赞数据
		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video := new(Video)
		if err := tx.WithContext(ctx).First(video, vid).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideos").Append(video); err != nil {
			return err
		}

		//2.改变 video 表中的 favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		//ErrDatabase在pkg/constants/error_const.go中定义
		//后续需要修改，因为需要返回的是数字而不是错误码，目前还没有理顺这个
		if res.RowsAffected != 1 {
			return constants.ErrDatabase
		}

		return nil
	})
	return err
}

// DisFavorite deletes the specified favorite from the database
func DisFavorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//1. 删除点赞数据
		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video, err := GetFavoriteRelation(ctx, uid, vid)
		if err != nil {
			return err
		}

		err = tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideos").Delete(video)
		if err != nil {
			return err
		}

		//2.改变 video 表中的 favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return constants.ErrDatabase
		}

		return nil
	})
	return err
}

// FavoriteList returns a list of Favorite videos.
func FavoriteList(ctx context.Context, uid int64) ([]Video, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	videos := []Video{}
	// if err := DB.WithContext(ctx).First(&video, vid).Error; err != nil {
	// 	return nil, err
	// }

	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&videos); err != nil {
		return nil, err
	}
	return videos, nil
}
