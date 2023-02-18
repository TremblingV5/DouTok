package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/service"
	"time"

	"github.com/TremblingV5/DouTok/applications/comment/dal/model"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func PackCommentListResp(code int32, msg string, comments []*model.CommentInHB) (resp *comment.DouyinCommentListResponse, err error) {
	resp = &comment.DouyinCommentListResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	comment_list := []*comment.Comment{}

	currentTime := time.Now().Unix()

	for _, v := range comments {
		temp := comment.Comment{
			Id:         v.GetId(),
			Content:    v.GetContent(),
			CreateDate: service.GetTimeRecall(v.GetTimestamp(), currentTime),
		}

		reqUser, err := rpc.GetUserById(context.Background(), &user.DouyinUserRequest{
			UserId: v.GetUserId(),
		})
		if err != nil {
			continue
		}

		tempUser := user.User{
			Id:              reqUser.User.Id,
			Name:            reqUser.User.Name,
			FollowCount:     reqUser.User.FollowCount,
			FollowerCount:   reqUser.User.FollowerCount,
			Avatar:          reqUser.User.Avatar,
			BackgroundImage: reqUser.User.BackgroundImage,
			Signature:       reqUser.User.Signature,
		}

		temp.User = &tempUser
		comment_list = append(comment_list, &temp)
	}

	resp.CommentList = comment_list

	return resp, nil
}
