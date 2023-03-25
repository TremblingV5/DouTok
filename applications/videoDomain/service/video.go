package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/model"
)

type GetVideoService struct {
	ctx context.Context
}

func NewGetVideoService(ctx context.Context) *GetVideoService {
	return &GetVideoService{ctx: ctx}
}

func (s *GetVideoService) GetVideoByIdInRDB(userId uint64) (*model.Video, error) {
	v, err := Do.Where(
		Video.ID.Eq(userId),
	).First()

	if err != nil {
		return v, err
	}

	return v, nil
}

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
