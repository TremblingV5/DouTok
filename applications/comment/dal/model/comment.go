package model

import "time"

type Comment struct {
	Id             int64
	VideoId        int64
	UserId         int64 // 留下评论的id
	ConversationId int64
	LastId         int64 // 二级回复所使用到的上一条回复的id
	ToUserId       int64 // 二级回复的用户id
	Content        string
	Timestamp      string
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
