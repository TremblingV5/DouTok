package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func PackCommentActionResp(code int32, msg string, cmt *comment.Comment, user_id int64) (resp *comment.DouyinCommentActionResponse, err error) {
	if code == int32(misc.Success.ErrCode) {
		reqUser, err := rpc.GetUserById(context.Background(), &user.DouyinUserRequest{
			UserId: user_id,
		})
		if err != nil {
			return nil, err
		}

		user := user.User{
			Id:              reqUser.User.Id,
			Name:            reqUser.User.Name,
			Avatar:          reqUser.User.Avatar,
			BackgroundImage: reqUser.User.BackgroundImage,
			Signature:       reqUser.User.Signature,
			FollowCount:     reqUser.User.FollowCount,
			FollowerCount:   reqUser.User.FollowerCount,
		}

		cmt.User = &user
	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Comment:    cmt,
	}, nil
}
