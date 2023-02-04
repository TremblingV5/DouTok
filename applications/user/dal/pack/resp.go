package pack

import (
	"errors"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

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

// BuilduserResp build userResp from error
func BuilduserUserResp(err error) *user.DouyinUserResponse {
	if err == nil {
		return userResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return userResp(s)
}

func userResp(err errno.ErrNo) *user.DouyinUserResponse {
	return &user.DouyinUserResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildVideoResp build VideoResp from error
func BuildVideoResp(err error) *feed.DouyinFeedResponse {
	if err == nil {
		return videoResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return videoResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return videoResp(s)
}

func videoResp(err errno.ErrNo) *feed.DouyinFeedResponse {
	return &feed.DouyinFeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildPublishResp build PublishResp from error
func BuildPublishResp(err error) *publish.DouyinPublishActionResponse {
	if err == nil {
		return publishResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return publishResp(s)
}

func publishResp(err errno.ErrNo) *publish.DouyinPublishActionResponse {
	return &publish.DouyinPublishActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildPublishResp build PublishResp from error
func BuildPublishListResp(err error) *publish.DouyinPublishListResponse {
	if err == nil {
		return publishListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return publishListResp(s)
}

func publishListResp(err errno.ErrNo) *publish.DouyinPublishListResponse {
	return &publish.DouyinPublishListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildFavoriteActionResp build FavoriteActionResp from error
func BuildFavoriteActionResp(err error) *favorite.DouyinFavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteActionResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return favoriteActionResp(s)
}

func favoriteActionResp(err errno.ErrNo) *favorite.DouyinFavoriteActionResponse {
	return &favorite.DouyinFavoriteActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildFavoriteListResp build FavoriteListResp from error
func BuildFavoriteListResp(err error) *favorite.DouyinFavoriteListResponse {
	if err == nil {
		return favoriteListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return favoriteListResp(s)
}

func favoriteListResp(err errno.ErrNo) *favorite.DouyinFavoriteListResponse {
	return &favorite.DouyinFavoriteListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildRelationActionResp build RelationActionResp from error
func BuildRelationActionResp(err error) *relation.DouyinRelationActionResponse {
	if err == nil {
		return relationActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationActionResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return relationActionResp(s)
}

func relationActionResp(err errno.ErrNo) *relation.DouyinRelationActionResponse {
	return &relation.DouyinRelationActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildFollowingListResp build FollowingListResp from error
func BuildFollowingListResp(err error) *relation.DouyinRelationFollowListResponse {
	if err == nil {
		return followingListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return followingListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return followingListResp(s)
}

func followingListResp(err errno.ErrNo) *relation.DouyinRelationFollowListResponse {
	return &relation.DouyinRelationFollowListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildFollowerListResp build FollowerListResp from error
func BuildFollowerListResp(err error) *relation.DouyinRelationFollowerListResponse {
	if err == nil {
		return followerListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return followerListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return followerListResp(s)
}

func followerListResp(err errno.ErrNo) *relation.DouyinRelationFollowerListResponse {
	return &relation.DouyinRelationFollowerListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildCommentActionResp build CommentActionResp from error
func BuildCommentActionResp(err error) *comment.DouyinCommentActionResponse {
	if err == nil {
		return commentActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentActionResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return commentActionResp(s)
}

func commentActionResp(err errno.ErrNo) *comment.DouyinCommentActionResponse {
	return &comment.DouyinCommentActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// BuildCommentListResp build CommentListResp from error
func BuildCommentListResp(err error) *comment.DouyinCommentListResponse {
	if err == nil {
		return commentListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return commentListResp(s)
}

func commentListResp(err errno.ErrNo) *comment.DouyinCommentListResponse {
	return &comment.DouyinCommentListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
