package model

import "encoding/binary"

type CommentInHB struct {
	Id             []byte `json:"id"`
	VideoId        []byte `json:"video_id"`
	UserId         []byte `json:"user_id"`
	ConversationId []byte `json:"conversation_id"`
	LastId         []byte `json:"last_id"`
	ToUserId       []byte `json:"to_user_id"`
	Content        []byte `json:"content"`
	Timestamp      []byte `json:"timestamp"`
}

func ToInt64(data []byte) int64 {
	return int64(binary.BigEndian.Uint64(data))
}

func (c *CommentInHB) GetId() int64 {
	return ToInt64(c.Id)
}

func (c *CommentInHB) GetVideoId() int64 {
	return ToInt64(c.VideoId)
}

func (c *CommentInHB) GetUserId() int64 {
	return ToInt64(c.UserId)
}

func (c *CommentInHB) GetConversationId() int64 {
	return ToInt64(c.ConversationId)
}

func (c *CommentInHB) GetLastId() int64 {
	return ToInt64(c.LastId)
}

func (c *CommentInHB) GetToUserId() int64 {
	return ToInt64(c.ToUserId)
}

func (c *CommentInHB) GetContent() string {
	return string(c.Content)
}

func (c *CommentInHB) GetTimestamp() string {
	return string(c.Timestamp)
}
