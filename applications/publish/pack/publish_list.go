package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"

	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
)

func PackagePublishListResponse(ctx context.Context, result *videoDomain.DoutokListPublishResponse, err error) (*publish.DouyinPublishListResponse, error) {
	if err != nil {
		return nil, err
	}

	if len(result.VideoList) <= 0 {
		return &publish.DouyinPublishListResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
			VideoList:  []*entity.Video{},
		}, nil
	}

	var userIdList []int64
	var videoIdList []int64
	for _, v := range result.VideoList {
		userIdList = append(userIdList, v.Author.GetId())
		videoIdList = append(videoIdList, v.GetId())
	}

	userInfo, err := rpc.UserDomainClient.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{
		UserId: userIdList,
	})
	if err != nil {
		return nil, err
	}

	favCount, err := rpc.FavoriteDomainClient.CountFavorite(ctx, &favoriteDomain.DoutokCountFavRequest{
		UserIdList: videoIdList,
	})
	if err != nil {
		return nil, err
	}

	var videoList []*entity.Video
	for _, v := range result.VideoList {
		temp := &entity.Video{
			Id:       v.GetId(),
			PlayUrl:  v.GetPlayUrl(),
			CoverUrl: v.GetCoverUrl(),
		}

		if u, ok := userInfo.UserList[v.Author.GetId()]; ok {
			temp.Author = &entity.User{
				Id:              u.Id,
				Name:            u.Name,
				Avatar:          u.Avatar,
				BackgroundImage: u.BackgroundImage,
				Signature:       u.Signature,
			}
		} else {
			temp.Author = &entity.User{
				Id: u.GetId(),
			}
		}

		if favCnt, ok := favCount.CountFav[v.Id]; ok {
			temp.FavoriteCount = favCnt
		}

		videoList = append(videoList, temp)
	}

	return &publish.DouyinPublishListResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		VideoList:  videoList,
	}, nil
}
