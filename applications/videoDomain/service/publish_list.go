package service

import (
	"context"
	"fmt"

	tools "github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/misc"
)

type QueryPublishListService struct {
	ctx context.Context
}

func NewQueryPublishListService(ctx context.Context) *QueryPublishListService {
	return &QueryPublishListService{ctx: ctx}
}

func (s *QueryPublishListService) QueryPublishListInHBase(userId int64) ([]*typedef.VideoInHB, error) {
	userIdString := tools.FillUserId(fmt.Sprint(userId))

	filters := hbaseHandle.GetFilterByRowKeyPrefix(userIdString)

	videoList, err := HBClient.Scan(
		"publish", filters...,
	)

	var list []*typedef.VideoInHB
	if err != nil {
		return list, err
	}

	for _, v := range videoList {
		temp := typedef.VideoInHB{}
		err := misc.Map2Struct4HB(v, &temp)
		if err != nil {
			continue
		}
		list = append(list, &temp)
	}

	return list, nil
}
