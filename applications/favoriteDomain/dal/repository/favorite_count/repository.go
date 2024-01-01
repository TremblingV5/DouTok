package favorite_count

import (
	"errors"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/query"
	"gorm.io/gorm"
)

type Repository interface {
	Load(videoId int64) (int64, error)
	Update(videoId, count int64) error
}

type PersistRepository struct {
	FavoriteCount query.IFavoriteCountDo
}

func (r *PersistRepository) Load(videoId int64) (int64, error) {
	count, err := r.FavoriteCount.Where(query.FavoriteCount.VideoId.Eq(videoId)).First()
	if err != nil {
		return 0, err
	}

	return count.Number, nil
}

func (r *PersistRepository) Update(videoId, count int64) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		number, err := tx.FavoriteCount.Where(query.FavoriteCount.VideoId.Eq(videoId)).First()
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := tx.FavoriteCount.Create(&model.FavoriteCount{
				VideoId: videoId,
				Number:  count,
			}); err != nil {
				return err
			}
		}

		_, err = tx.FavoriteCount.Updates(&model.FavoriteCount{
			VideoId: videoId,
			Number:  number.Number + count,
		})
		return err
	})
}
