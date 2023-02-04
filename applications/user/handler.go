package main

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/applications/user/command"
	"github.com/TremblingV5/DouTok/applications/user/dal/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/TremblingV5/DouTok/pkg/jwt"
)

// UserSrvImpl implements the user service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserSrvImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {

		resp = pack.BuilduserRegisterResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewCreateUserService(ctx).CreateUser(req, Argon2Config)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	// 新用户注册成功后直接登录
	uid, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(uid),
	})
	if err != nil {
		resp = pack.BuilduserRegisterResp(errno.ErrSignatureInvalid)
		return resp, nil
	}

	resp = pack.BuilduserRegisterResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// Login implements the UserSrvImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuilduserRegisterResp(errno.ErrBind)
		return resp, nil
	}

	uid, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(uid),
	})
	if err != nil {
		resp = pack.BuilduserRegisterResp(errno.ErrSignatureInvalid)
		return resp, nil
	}

	resp = pack.BuilduserRegisterResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// GetUserById implements the UserSrvImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuilduserUserResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	// else if claim.Id != int64(req.UserId) {
	// 	resp = pack.BuilduserUserResp(errno.ErrValidation)
	// 	return resp, nil
	// }

	if req.UserId < 0 {
		resp = pack.BuilduserUserResp(errno.ErrBind)
		return resp, nil
	}

	user, err := command.NewMGetUserService(ctx).MGetUser(req, claim.Id)
	if err != nil {
		resp = pack.BuilduserUserResp(err)
		return resp, nil
	}

	//if claim.Id == req.UserId {
	//	user.IsFollow = true
	//} else {
	//	// TODO 获取claim.id 是否已关注 req.userid
	//	user.IsFollow = false
	//}

	resp = pack.BuilduserUserResp(errno.Success)
	resp.User = user
	return resp, nil
}

// BuilduserRegisterResp build userRegisterResp from error

func BuilduserRegisterResp(err error) *user.DouyinUserRegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return userRegisterResp(s)
}
func userRegisterResp(err errno.ErrNo) *user.DouyinUserRegisterResponse {
	return &user.DouyinUserRegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

type ErrNo struct {
	ErrCode int
	ErrMsg  string
}
