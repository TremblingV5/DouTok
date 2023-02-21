package rpc

import (
	"context"
	"testing"

	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	mock "github.com/TremblingV5/DouTok/pkg/mock/comment"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/golang/mock/gomock"
)

func TestCommentList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	client.EXPECT().CommentList(gomock.Eq(context.Background()), gomock.Eq(&comment.DouyinCommentListRequest{VideoId: 666666})).
		Return(&comment.DouyinCommentListResponse{
			CommentList: []*comment.Comment{
				{
					Id: int64(10001000),
					User: &user.User{
						Id:              int64(10001000),
						Name:            "integration",
						FollowerCount:   int64(10),
						FollowCount:     int64(10),
						IsFollow:        true,
						Avatar:          "integration",
						BackgroundImage: "integration",
						Signature:       "integration",
						TotalFavorited:  int64(10),
						WorkCount:       int64(10),
						FavoriteCount:   int64(10),
					},
					Content:    "integration",
					LikeCount:  int64(10),
					TeaseCount: int64(10),
					CreateDate: "integration",
				},
			},
		}, nil)

	resp, err := CommentList(context.Background(), client, &comment.DouyinCommentListRequest{VideoId: 666666})
	assert.DeepEqual(t, int64(10001000), resp.CommentList[0].User.Id)
	assert.DeepEqual(t, int64(10), resp.CommentList[0].User.FollowCount)
	assert.DeepEqual(t, int64(10), resp.CommentList[0].User.FollowerCount)
	assert.DeepEqual(t, int64(10), resp.CommentList[0].User.TotalFavorited)
	assert.DeepEqual(t, int64(10), resp.CommentList[0].User.WorkCount)
	assert.DeepEqual(t, int64(10), resp.CommentList[0].User.FavoriteCount)
	assert.DeepEqual(t, true, resp.CommentList[0].User.IsFollow)
	assert.DeepEqual(t, "integration", resp.CommentList[0].User.Name)
	assert.DeepEqual(t, "integration", resp.CommentList[0].User.Avatar)
	assert.DeepEqual(t, "integration", resp.CommentList[0].User.BackgroundImage)
	assert.DeepEqual(t, "integration", resp.CommentList[0].Content)
	assert.DeepEqual(t, int64(10), resp.CommentList[0].LikeCount)
	assert.DeepEqual(t, int64(10), resp.CommentList[0].TeaseCount)
	assert.DeepEqual(t, "integration", resp.CommentList[0].CreateDate)
	assert.Nil(t, err)
}

func TestCommentAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	client.EXPECT().CommentAction(gomock.Eq(context.Background()), gomock.Eq(&comment.DouyinCommentActionRequest{VideoId: 666666})).
		Return(&comment.DouyinCommentActionResponse{
			Comment: &comment.Comment{
				Id: int64(10001000),
				User: &user.User{
					Id:              int64(10001000),
					Name:            "integration",
					FollowerCount:   int64(10),
					FollowCount:     int64(10),
					IsFollow:        true,
					Avatar:          "integration",
					BackgroundImage: "integration",
					Signature:       "integration",
					TotalFavorited:  int64(10),
					WorkCount:       int64(10),
					FavoriteCount:   int64(10),
				},
				Content:    "integration",
				LikeCount:  int64(10),
				TeaseCount: int64(10),
				CreateDate: "integration",
			},
		}, nil)

	resp, err := CommentAction(context.Background(), client, &comment.DouyinCommentActionRequest{VideoId: 666666})
	assert.DeepEqual(t, int64(10001000), resp.Comment.User.Id)
	assert.DeepEqual(t, int64(10), resp.Comment.User.FollowCount)
	assert.DeepEqual(t, int64(10), resp.Comment.User.FollowerCount)
	assert.DeepEqual(t, int64(10), resp.Comment.User.TotalFavorited)
	assert.DeepEqual(t, int64(10), resp.Comment.User.WorkCount)
	assert.DeepEqual(t, int64(10), resp.Comment.User.FavoriteCount)
	assert.DeepEqual(t, true, resp.Comment.User.IsFollow)
	assert.DeepEqual(t, "integration", resp.Comment.User.Name)
	assert.DeepEqual(t, "integration", resp.Comment.User.Avatar)
	assert.DeepEqual(t, "integration", resp.Comment.User.BackgroundImage)
	assert.DeepEqual(t, "integration", resp.Comment.Content)
	assert.DeepEqual(t, int64(10), resp.Comment.LikeCount)
	assert.DeepEqual(t, int64(10), resp.Comment.TeaseCount)
	assert.DeepEqual(t, "integration", resp.Comment.CreateDate)
	assert.Nil(t, err)
}
