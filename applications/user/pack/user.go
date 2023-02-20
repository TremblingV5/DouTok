package pack

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/dal/model"
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func PackUserResp(code int32, msg string, u *model.User) (resp *user.DouyinUserResponse, err error) {
	resp = &user.DouyinUserResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	userResp, err := rpc.GetFollowCount(context.Background(), &relation.DouyinRelationCountRequest{
		UserId: int64(u.ID),
	})

	if err != nil {
		resp.StatusCode = int32(misc.SystemErr.ErrCode)
		resp.StatusMsg = misc.SystemErr.ErrMsg
		return resp, err
	}

	info := user.User{
		Id:              int64(u.ID),
		Name:            u.UserName,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		FollowCount:     userResp.FollowCount,
		FollowerCount:   userResp.FollowerCount,
		IsFollow:        true,
	}

	resp.User = &info

	return resp, nil
}
