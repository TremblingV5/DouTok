package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func PackageFeedListResp(ctx context.Context, result *videoDomain.DoutokGetFeedResponse, userId int64, err error) (*feed.DouyinFeedResponse, error) {
	if err != nil {
		return nil, err
	}

	var videoIdList []int64
	var userIdList []int64
	for _, v := range result.VideoList {
		videoIdList = append(videoIdList, v.Id)
		userIdList = append(userIdList, v.Author.Id)
	}

	userInfo, err := rpc.UserDomainClient.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{
		UserId: userIdList,
	})
	if err != nil {
		return nil, err
	}

	isFavInfo, err := rpc.FavoriteDomainClient.IsFavorite(ctx, &favoriteDomain.DoutokIsFavRequest{
		UserId:  userId,
		VideoId: videoIdList,
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

	commentCount, err := rpc.CommentDomainClient.CountComment(ctx, &commentDomain.DoutokCountCommentReq{
		VideoIdList: videoIdList,
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
			Title:    v.GetTitle(),
		}

		if u, ok := userInfo.UserList[v.GetId()]; ok {
			temp.Author = &entity.User{
				Id:              u.Id,
				Name:            u.Name,
				Avatar:          u.Avatar,
				BackgroundImage: u.BackgroundImage,
				Signature:       u.Signature,
			}
		} else {
			temp.Author = &entity.User{
				Id: v.Author.Id,
			}
		}

		if isFav, ok := isFavInfo.IsFav[v.GetId()]; ok {
			temp.IsFavorite = isFav
		} else {
			temp.IsFavorite = false
		}

		if favCnt, ok := favCount.CountFav[v.GetId()]; ok {
			temp.FavoriteCount = favCnt
		} else {
			temp.FavoriteCount = 0
		}

		if commentCnt, ok := commentCount.CommentCount[v.GetId()]; ok {
			temp.CommentCount = commentCnt
		} else {
			temp.CommentCount = 0
		}

		videoList = append(videoList, temp)
	}

	return &feed.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "Success",
		VideoList:  videoList,
		NextTime:   result.NextTime,
	}, nil
}
