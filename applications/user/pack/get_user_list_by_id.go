package pack

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/dal/model"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func PackUserListResp(code int32, msg string, list []*model.User) (resp *user.DouyinUserListResponse, err error) {
	resp = &user.DouyinUserListResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	result := []*user.User{}

	for _, v := range list {
		req, err := rpc.GetFollowCount(
			context.Background(),
			&relation.DouyinRelationCountRequest{UserId: int64(v.ID)},
		)
		if err != nil {
			continue
		}

		temp := &user.User{
			Id:              int64(v.ID),
			Name:            v.UserName,
			Avatar:          v.Avatar,
			BackgroundImage: v.BackgroundImage,
			Signature:       v.Signature,
			FollowCount:     req.FollowCount,
			FollowerCount:   req.FollowerCount,
			IsFollow:        false,
		}

		result = append(result, temp)
	}

	resp.UserList = result
	return resp, nil
}
