package comment_api

import (
	"context"
	commententity "github.com/TremblingV5/DouTok/applications/comment/domain/entity/comment"
	"github.com/TremblingV5/DouTok/applications/comment/infra/misc"
	"github.com/TremblingV5/DouTok/applications/comment/services/comment_service"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"time"
)

func (s *CommentServiceImpl) packCommentActionResp(code int32, msg string, cmt *commententity.Entity, userId int64) (resp *comment.DouyinCommentActionResponse, err error) {
	var result *comment.Comment

	result.Id = cmt.Id
	result.Content = cmt.Content
	result.CreateDate = cmt.Timestamp

	if code == int32(misc.Success.ErrCode) {
		reqUser, err := s.clients.User.Client.GetUserById(context.Background(), &user.DouyinUserRequest{
			UserId: userId,
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

		result.User = &user
	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Comment:    result,
	}, nil
}

func (s *CommentServiceImpl) packCommentCountResp(code int32, msg string, countList map[int64]int64) (*comment.DouyinCommentCountResponse, error) {
	return &comment.DouyinCommentCountResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Result:     countList,
	}, nil
}

func (s *CommentServiceImpl) packCommentListResp(code int32, msg string, comments []*commententity.Entity) (resp *comment.DouyinCommentListResponse, err error) {
	resp = &comment.DouyinCommentListResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	var commentList []*comment.Comment

	currentTime := time.Now().Unix()

	for _, v := range comments {
		temp := comment.Comment{
			Id:         v.Id,
			Content:    v.Content,
			CreateDate: comment_service.GetTimeRecall(v.Timestamp, currentTime),
		}

		reqUser, err := s.clients.User.Client.GetUserById(context.Background(), &user.DouyinUserRequest{
			UserId: v.UserId,
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
		commentList = append(commentList, &temp)
	}

	resp.CommentList = commentList

	return resp, nil
}
