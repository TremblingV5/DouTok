package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func PackageCommentListRepsonse(ctx context.Context, result *commentDomain.DoutokListCommentResp, e error) (resp *comment.DouyinCommentListResponse, err error) {
	if e != nil {
		return nil, e
	}

	var userIdList []int64
	for _, v := range result.CommentList {
		userIdList = append(userIdList, v.User.GetId())
	}

	userInfo, err := rpc.UserDomainRPCClient.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{
		UserId: userIdList,
	})

	var commentList []*entity.Comment
	for _, v := range result.CommentList {
		temp := &entity.Comment{
			Id:         v.GetId(),
			Content:    v.GetContent(),
			CreateDate: v.GetCreateDate(),
		}

		if v, ok := userInfo.UserList[v.User.GetId()]; ok {
			temp.User = &entity.User{
				Id:              v.GetId(),
				Name:            v.GetName(),
				Avatar:          v.GetAvatar(),
				BackgroundImage: v.GetBackgroundImage(),
				Signature:       v.GetSignature(),
			}
		} else {
			temp.User = &entity.User{
				Id: v.GetId(),
			}
		}
		commentList = append(commentList, temp)
	}

	return &comment.DouyinCommentListResponse{
		StatusCode:  result.StatusCode,
		StatusMsg:   result.StatusMsg,
		CommentList: commentList,
	}, e
}
