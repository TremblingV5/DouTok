package db
//目前只改了import和结构体
import (
	"context"
	"gorm.io/gorm"
	"github.com/TremblingV5/DouTok/pkg/constants"
)

type UserFavorVideos struct {
	gorm.Model
	UserId  int64  `json:"user_id"`
	VideoId  int64  `json:"video_id"`
}

func (f *UserFavorVideos) TableName() string {
	return constants.FavoriteTableName
}

// CreateFavorite
func CreateFavorite(ctx context.Context, favors []*UserFavorVideos) error {
	if err := DB.WithContext(ctx).Create(favors).Error; err != nil {
		return err
	}
	return nil
}

//  get list of favor info
func MGetNotes(ctx context.Context, noteIDs []int64) ([]*UserFavorVideos, error) {
	var res []*UserFavorVideos
	if len(noteIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", noteIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// UpdateNote update note info
func UpdateNote(ctx context.Context, noteID, userID int64, title, content *string) error {
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if content != nil {
		params["content"] = *content
	}
	return DB.WithContext(ctx).Model(&UserFavorVideos{}).Where("id = ? and user_id = ?", noteID, userID).
		Updates(params).Error
}

// DeleteNote delete note info
func DeleteNote(ctx context.Context, noteID, userID int64) error {
	return DB.WithContext(ctx).Where("id = ? and user_id = ? ", noteID, userID).Delete(&UserFavorVideos{}).Error
}

// QueryNote query list of note info
func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*UserFavorVideos, int64, error) {
	var total int64
	var res []*UserFavorVideos
	conn := DB.WithContext(ctx).Model(&UserFavorVideos{}).Where("user_id = ?", userID)

	if searchKey != nil {
		conn = conn.Where("title like ?", "%"+*searchKey+"%")
	}

	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}

	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}
