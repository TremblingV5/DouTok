package service

import (
	"github.com/TremblingV5/DouTok/applications/publish/dal/model"
)

func QueryVideoFromRBDById(id uint64) (*model.Video, error) {
	v, err := Do.Where(
		Video.ID.Eq(id),
	).First()

	if err != nil {
		return v, err
	}

	return v, nil
}

func QuerySomeVideoFromRDBByIds(id ...uint64) ([]*model.Video, error) {
	videos, err := Do.Where(
		Video.ID.In(id...),
	).Find()

	if err != nil {
		return videos, err
	}

	return videos, nil
}
