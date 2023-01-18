package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/feed/dal/model"
	"github.com/TremblingV5/DouTok/applications/feed/dal/query"
)

func QueryVideoCountFromRDB(id uint64) (*model.Video, error) {
	query.SetDefault(DB)
	video := query.Video
	do := video.WithContext(context.Background())

	v, err := do.Where(
		video.ID.Eq(id),
	).First()

	if err != nil {
		return v, err
	}

	return v, nil
}

func QueryAFewVideoCountFromRDB(id ...uint64) ([]*model.Video, error) {
	query.SetDefault(DB)
	video := query.Video
	do := video.WithContext(context.Background())

	videos, err := do.Where(
		video.ID.In(id...),
	).Find()

	if err != nil {
		return videos, err
	}

	return videos, nil
}
