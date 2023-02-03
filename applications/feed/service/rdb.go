package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/feed/dal/model"
)

func GetVideoByIdInRDB(ctx context.Context, user_id uint64) (*model.Video, error) {
	v, err := Do.Where(
		Video.ID.Eq(user_id),
	).First()

	if err != nil {
		return v, err
	}

	return v, nil
}
