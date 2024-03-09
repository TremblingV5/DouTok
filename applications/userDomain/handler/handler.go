package handler

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/userDomain/errs"
	"github.com/TremblingV5/DouTok/applications/userDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

type Handler struct {
	user *service.Service
}

func New(user *service.Service) *Handler {
	return &Handler{
		user: user,
	}
}

func (s *Handler) AddUser(ctx context.Context, req *userDomain.DoutokAddUserRequest) (resp *userDomain.DoutokAddUserResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return &userDomain.DoutokAddUserResponse{
			StatusCode: errs.EmptyErr.Code(),
			StatusMsg:  errs.EmptyErr.Message(),
			UserId:     0,
		}, nil
	}

	userId, err := s.user.CreateNewUser(req.Username, req.Password)

	if err != nil {
		return &userDomain.DoutokAddUserResponse{
			StatusCode: errs.SystemErr.Code(),
			StatusMsg:  fmt.Sprintf("%s %s", errs.SystemErr.Message(), err.Error()),
			UserId:     userId,
		}, err
	}

	return &userDomain.DoutokAddUserResponse{
		StatusCode: errs.Success.Code(),
		StatusMsg:  errs.Success.Message(),
		UserId:     userId,
	}, nil
}

func (s *Handler) CheckUser(ctx context.Context, req *userDomain.DoutokCheckUserRequest) (resp *userDomain.DoutokCheckUserResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return packCheckUserResp(errs.EmptyErr.Code(), errs.EmptyErr.Message(), 0)
	}

	userId, err := s.user.CheckPassword(req.Username, req.Password)
	if err != nil || userId == 0 {
		return packCheckUserResp(errs.PasswordErr.Code(), errs.PasswordErr.Message(), 0)
	}

	return packCheckUserResp(errs.Success.Code(), errs.Success.Message(), userId)
}

func (s *Handler) GetUserInfo(ctx context.Context, req *userDomain.DoutokGetUserInfoRequest) (resp *userDomain.DoutokGetUserInfoResponse, err error) {
	userIdList := make([]uint64, 0)
	if req != nil && req.UserId != nil {
		for _, v := range req.UserId {
			userIdList = append(userIdList, uint64(v))
		}
	}

	userList, err := s.user.LoadUserListByIds(userIdList...)
	if err != nil {
		return packGetUserInfoResp(errs.SystemErr.Code(), err.Error(), nil)
	}

	return packGetUserInfoResp(errs.Success.Code(), errs.Success.Message(), userList)
}
