package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func (s *VideoDomainServiceImpl) GetVideoInfo(ctx context.Context, req *videoDomain.DoutokGetVideoInfoRequest) (resp *entity.Video, err error) {
	data, err := service.NewGetVideoService(ctx).GetVideoByIdInRDB(uint64(req.VideoId))
	if err != nil {
		return nil, err
	}
	return &entity.Video{
		Id: int64(data.ID),
		Author: &entity.User{
			Id: int64(data.AuthorID),
		},
		PlayUrl:  data.VideoUrl,
		CoverUrl: data.CoverUrl,
		Title:    data.Title,
	}, nil
}
