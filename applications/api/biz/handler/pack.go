package handler

import (
	"errors"
	"github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}

// message
func messageChatResp(err errno.ErrNo) *api.DouyinMessageChatResponse {
	return &api.DouyinMessageChatResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func messageActionResp(err errno.ErrNo) *api.DouyinMessageActionResponse {
	return &api.DouyinMessageActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildMessageChatResp(err error) *api.DouyinMessageChatResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return messageChatResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return messageChatResp(e)
}

func BuildMessageActionResp(err error) *api.DouyinMessageActionResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return messageActionResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return messageActionResp(e)
}

// user
func userRegisterResp(err errno.ErrNo) *api.DouyinUserRegisterResponse {
	return &api.DouyinUserRegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func userLoginResp(err errno.ErrNo) *api.DouyinUserLoginResponse {
	return &api.DouyinUserLoginResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func getUserResp(err errno.ErrNo) *api.DouyinUserResponse {
	return &api.DouyinUserResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildUserRegisterResp(err error) *api.DouyinUserRegisterResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return userRegisterResp(e)
}

func BuildUserLoginResp(err error) *api.DouyinUserLoginResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userLoginResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return userLoginResp(e)
}

func BuildGetUserResp(err error) *api.DouyinUserResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getUserResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return getUserResp(e)
}

// relation
func relationActionResp(err errno.ErrNo) *api.DouyinRelationActionResponse {
	return &api.DouyinRelationActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func relationFollowListResp(err errno.ErrNo) *api.DouyinRelationFollowListResponse {
	return &api.DouyinRelationFollowListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func relationFollowerListResp(err errno.ErrNo) *api.DouyinRelationFollowerListResponse {
	return &api.DouyinRelationFollowerListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func relationFriendListResp(err errno.ErrNo) *api.DouyinRelationFriendListResponse {
	return &api.DouyinRelationFriendListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildRelationActionResp(err error) *api.DouyinRelationActionResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationActionResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return relationActionResp(e)
}

func BuildRelationFollowListResp(err error) *api.DouyinRelationFollowListResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFollowListResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return relationFollowListResp(e)
}

func BuildRelationFollowerListResp(err error) *api.DouyinRelationFollowerListResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFollowerListResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return relationFollowerListResp(e)
}

func BuildRelationFriendListResp(err error) *api.DouyinRelationFriendListResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFriendListResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return relationFriendListResp(e)
}

// publish
func publishActionResp(err errno.ErrNo) *api.DouyinPublishActionResponse {
	return &api.DouyinPublishActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func publishListResp(err errno.ErrNo) *api.DouyinPublishListResponse {
	return &api.DouyinPublishListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildPublishActionResp(err error) *api.DouyinPublishActionResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishActionResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return publishActionResp(e)
}

func BuildPublishListResp(err error) *api.DouyinPublishListResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishListResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return publishListResp(e)
}

// feed
func getUserFeedResp(err errno.ErrNo) *api.DouyinFeedResponse {
	return &api.DouyinFeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildGetUserFeedResp(err error) *api.DouyinFeedResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getUserFeedResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return getUserFeedResp(e)
}

// favorite
func favoriteActionResp(err errno.ErrNo) *api.DouyinFavoriteActionResponse {
	return &api.DouyinFavoriteActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func favoriteListResp(err errno.ErrNo) *api.DouyinFavoriteListResponse {
	return &api.DouyinFavoriteListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildFavoriteActionResp(err error) *api.DouyinFavoriteActionResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteActionResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return favoriteActionResp(e)
}

func BuildFavoriteListResp(err error) *api.DouyinFavoriteListResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteListResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return favoriteListResp(e)
}

// comment
func commentActionResp(err errno.ErrNo) *api.DouyinCommentActionResponse {
	return &api.DouyinCommentActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func commentListResp(err errno.ErrNo) *api.DouyinCommentListResponse {
	return &api.DouyinCommentListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildCommendActionResp(err error) *api.DouyinCommentActionResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentActionResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return commentActionResp(e)
}

func BuildCommendListResp(err error) *api.DouyinCommentListResponse {
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListResp(e)
	}
	e = errno.InternalErr.WithMessage(err.Error())
	return commentListResp(e)
}
