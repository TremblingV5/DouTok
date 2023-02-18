package model

import "time"

type Comment struct {
	Id             int64 `gorm:"primarykey"`
	VideoId        int64 `gorm:"index:video_id"`
	UserId         int64 `gorm:"index:user_id"` // 留下评论的id
	ConversationId int64 `gorm:"index:con_id"`
	LastId         int64 // 二级回复所使用到的上一条回复的id
	ToUserId       int64 // 二级回复的用户id
	Content        string
	Timestamp      string
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
