package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/user/errs"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

type Handler struct {
	clients *rpc.Clients
}

func New() *Handler {
	return &Handler{
		clients: rpc.New(),
	}
}

func (s *Handler) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	result, err := s.clients.User.CheckUser(ctx, &userDomain.DoutokCheckUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &user.DouyinUserLoginResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		UserId:     result.UserId,
	}, nil
}

func (s *Handler) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	result, err := s.clients.User.AddUser(ctx, &userDomain.DoutokAddUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &user.DouyinUserRegisterResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		UserId:     result.UserId,
	}, nil
}

func (s *Handler) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	userInfo, err := s.clients.User.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{UserId: []int64{req.UserId}})
	if err != nil {
		return nil, err
	}

	resp = &user.DouyinUserResponse{
		StatusCode: userInfo.StatusCode,
		StatusMsg:  userInfo.StatusMsg,
	}

	if userInfo == nil || len(userInfo.UserList) <= 0 {
		resp.StatusCode = int32(errs.EmptyUserListErrCode)
		resp.StatusMsg = errs.EmptyUserListErr.ErrMsg
		return resp, nil
	}
	resp.User = userInfo.UserList[req.UserId]

	// hydrate follow count
	followCount, err := s.clients.Relation.CountRelation(ctx, &relationDomain.DoutokCountRelationRequest{
		UserId:     []int64{req.UserId},
		ActionType: int64(0),
	})
	if err != nil {
		resp.StatusCode = int32(errno.BadRequest.ErrCode)
		resp.StatusMsg = errno.BadRequest.ErrMsg + " relation domain rpc error"
		return resp, err
	}
	resp.User.FollowCount = followCount.Result[req.UserId]

	// hydrate follower count
	followerCount, err := s.clients.Relation.CountRelation(ctx, &relationDomain.DoutokCountRelationRequest{
		UserId:     []int64{req.UserId},
		ActionType: int64(1),
	})
	if err != nil {
		resp.StatusCode = int32(errno.BadRequest.ErrCode)
		resp.StatusMsg = errno.BadRequest.ErrMsg + " relation domain rpc error"
		return resp, err
	}
	resp.User.FollowerCount = followerCount.Result[req.UserId]

	return resp, nil
}
