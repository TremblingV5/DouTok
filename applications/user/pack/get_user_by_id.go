package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/errno"

	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func PackageGetUserByIdResponse(ctx context.Context, userInfo *userDomain.DoutokGetUserInfoResponse, userId int64, e error) (resp *user.DouyinUserResponse, err error) {
	if e != nil {
		return nil, e
	}

	resp = &user.DouyinUserResponse{
		StatusCode: userInfo.StatusCode,
		StatusMsg:  userInfo.StatusMsg,
	}

	if userInfo == nil || len(userInfo.UserList) <= 0 {
		resp.StatusCode = int32(misc.EmptyUserListErrCode)
		resp.StatusMsg = misc.EmptyUserListErr.ErrMsg
		return resp, e
	}

	if e != nil {
		return resp, e
	}

	resp.User = userInfo.UserList[userId]

	followCount, e := rpc.RelationDomainRPCClient.CountRelation(ctx, &relationDomain.DoutokCountRelationRequest{
		UserId:     []int64{userId},
		ActionType: int64(0),
	})
	if e != nil {
		resp.StatusCode = int32(errno.BadRequest.ErrCode)
		resp.StatusMsg = errno.BadRequest.ErrMsg + " relation domain rpc error"
		return resp, e
	}
	resp.User.FollowCount = followCount.Result[userId]

	followerCount, e := rpc.RelationDomainRPCClient.CountRelation(ctx, &relationDomain.DoutokCountRelationRequest{
		UserId:     []int64{userId},
		ActionType: int64(1),
	})
	if e != nil {
		resp.StatusCode = int32(errno.BadRequest.ErrCode)
		resp.StatusMsg = errno.BadRequest.ErrMsg + " relation domain rpc error"
		return resp, e
	}
	resp.User.FollowerCount = followerCount.Result[userId]

	return resp, e
}
