package service

import (
	"fmt"
	"time"

	"github.com/TremblingV5/DouTok/applications/comment/dal/model"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

/*
	添加评论
*/
func AddComment(video_id int64, user_id int64, con_id int64, last_id int64, to_user_id int64, content string) (*comment.Comment, error) {
	timestamp := fmt.Sprint(time.Now().Unix())

	id := utils.GetSnowFlakeId()

	result := comment.Comment{
		Id:         int64(id),
		Content:    content,
		CreateDate: "刚刚",
	}

	if err := SaveComment2RDB(int64(id), video_id, user_id, con_id, last_id, to_user_id, content, timestamp); err != nil {
		return nil, err
	}

	if err := SaveComment2HB(int64(id), video_id, user_id, con_id, last_id, to_user_id, content, timestamp); err != nil {
		return nil, err
	}

	UpdateCacheComCount(video_id, true)

	//reqUser, err := rpc.GetUserById(context.Background(), &user.DouyinUserRequest{
	//	UserId: user_id,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//user := user.User{
	//	Id:            reqUser.User.Id,
	//	Name:          reqUser.User.Name,
	//	Avatar:        reqUser.User.Avatar,
	//	FollowCount:   reqUser.User.FollowCount,
	//	FollowerCount: reqUser.User.FollowerCount,
	//}
	//
	//result.User = &user

	return &result, nil
}

/*
	在MySQL中存储评论信息
*/
func SaveComment2RDB(id int64, video_id int64, user_id int64, con_id int64, last_id int64, to_user_id int64, content string, timestamp string) error {
	err := DoComment.Create(
		&model.Comment{
			Id:             id,
			VideoId:        video_id,
			UserId:         user_id,
			ConversationId: con_id,
			LastId:         last_id,
			ToUserId:       to_user_id,
			Content:        content,
			Status:         true,
			Timestamp:      timestamp,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

/*
	在HBase中存储评论信息
*/
func SaveComment2HB(id int64, video_id int64, user_id int64, con_id int64, last_id int64, to_user_id int64, content string, timestamp string) error {
	hbData := map[string]map[string][]byte{
		"data": {
			"id":              []byte(fmt.Sprint(id)),
			"video_id":        []byte(fmt.Sprint(video_id)),
			"user_id":         []byte(fmt.Sprint(user_id)),
			"conversation_id": []byte(fmt.Sprint(con_id)),
			"last_id":         []byte(fmt.Sprint(last_id)),
			"to_user_id":      []byte(fmt.Sprint(to_user_id)),
			"content":         []byte(content),
			"timestamp":       []byte(timestamp),
		},
	}

	if err := HBClient.Put(
		"comment", misc.GetCommentRowKey(video_id, "0", con_id, timestamp), hbData,
	); err != nil {
		return err
	}

	return nil
}
