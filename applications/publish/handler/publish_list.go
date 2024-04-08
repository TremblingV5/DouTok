package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/publish/pack"

	"github.com/TremblingV5/DouTok/applications/publish/service"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
)

func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	user_id := req.UserId

	list, err := service.QueryPublishListInHBase(user_id)

	if err != nil {
		return nil, err
	}

	resp, err = pack.PackPublishListRes(list, 0, "Success", req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
