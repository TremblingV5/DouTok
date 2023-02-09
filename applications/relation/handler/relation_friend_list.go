package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
)

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	followerIds, err := db.GetFollowerIds(req.UserId)
	if err != nil {
		return nil, err
	}
	r, err := rpc.MessageClient.MessageFriendList(ctx, &message.DouyinFriendListMessageRequest{
		FriendIdList: followerIds,
		UserId:       req.UserId,
	})
	if err != nil {
		return nil, err
	}
	result := r.Result
	users, err := db.GetFollowList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userList := make([]*relation.FriendUser, len(followerIds))
	for i, u := range users {
		userList[i].User = u
		userList[i].Message = result[u.Id].Content
	}
	return &relation.DouyinRelationFriendListResponse{StatusMsg: "success", StatusCode: 0, UserList: userList}, nil
}
