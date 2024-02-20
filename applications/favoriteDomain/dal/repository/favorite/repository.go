package favorite

import (
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/query"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

type Repository interface {
	Load(userId, videoId int64) (bool, error)
	DoFavorite(userId, videoId int64) error
	UndoFavorite(userId, videoId int64) error
	CountFavoriteForUserId(userId int64) (int64, error)
	CountFavoriteForVideoId(videoId int64) (int64, error)
	LoadFavoriteListByUserId(userId int64) ([]int64, error)
}

type PersistRepository struct {
	favorite        query.IFavoriteDo
	snowflakeHandle *utils.SnowflakeHandle
}

func (r *PersistRepository) Load(userId, videoId int64) (*model.Favorite, error) {
	record, err := r.favorite.Where(query.Favorite.UserId.Eq(userId), query.Favorite.VideoId.Eq(videoId)).First()
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *PersistRepository) DoFavorite(userId, videoId int64) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		favorite, err := r.Load(userId, videoId)
		if err != nil {
			e := tx.Favorite.Create(&model.Favorite{ID: r.snowflakeHandle.GetId().Int64(), UserId: userId, VideoId: videoId, Status: 1})
			if err != nil {
				return e
			}
			return nil
		}
		_, err = tx.Favorite.Updates(&model.Favorite{ID: favorite.ID, Status: 1})
		return err
	})
}

func (r *PersistRepository) UndoFavorite(userId, videoId int64) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		favorite, err := r.Load(userId, videoId)
		if err != nil {
			e := tx.Favorite.Create(&model.Favorite{UserId: userId, VideoId: videoId, Status: 0})
			if err != nil {
				return e
			}
			return nil
		}
		_, err = tx.Favorite.Updates(&model.Favorite{ID: favorite.ID, Status: 0})
		return err
	})
}

func (r *PersistRepository) CountFavoriteForUserId(userId int64) (int64, error) {
	return r.favorite.Where(query.Favorite.UserId.Eq(userId)).Count()
}

func (r *PersistRepository) CountFavoriteForVideoId(videoId int64) (int64, error) {
	return r.favorite.Where(query.Favorite.VideoId.Eq(videoId)).Count()
}

func (r *PersistRepository) LoadFavoriteListByUserId(userId int64) ([]int64, error) {
	var list []int64
	err := r.favorite.Select(query.Favorite.ID).Where(query.Favorite.UserId.Eq(userId)).Scan(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
