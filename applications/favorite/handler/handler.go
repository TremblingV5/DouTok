package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/errs"
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

type Handler struct {
	client *rpc.Clients
}

func New(client *rpc.Clients) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	if req.ActionType == 1 {
		// 点赞
		result, err := h.client.Favorite.Client.AddFavorite(ctx, &favoriteDomain.DoutokAddFavRequest{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
		return &favorite.DouyinFavoriteActionResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
		}, err
	}

	if req.ActionType == 2 {
		// 取消点赞
		result, _ := h.client.Favorite.Client.RmFavorite(ctx, &favoriteDomain.DoutokRmFavRequest{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
		return &favorite.DouyinFavoriteActionResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
		}, err
	}

	return &favorite.DouyinFavoriteActionResponse{
		StatusCode: errs.BindingErr.Code(),
		StatusMsg:  errs.BindingErr.Message(),
	}, err
}

func (h *Handler) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	result, err := h.client.Favorite.Client.ListFavorite(ctx, &favoriteDomain.DoutokListFavRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return &favorite.DouyinFavoriteListResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
		}, err
	}

	var videoIdList []*entity.Video
	for _, v := range result.VideoList {
		videoInfo, e := h.client.Video.Client.GetVideoInfo(ctx, &videoDomain.DoutokGetVideoInfoRequest{
			VideoId: v.Id,
		})
		if e != nil {
			continue
		}
		videoIdList = append(videoIdList, videoInfo)
	}

	return &favorite.DouyinFavoriteListResponse{
		StatusCode: errs.Success.Code(),
		StatusMsg:  errs.Success.Message(),
		VideoList:  videoIdList,
	}, nil
}
