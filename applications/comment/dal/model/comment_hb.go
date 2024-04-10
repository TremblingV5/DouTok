package model

import (
	"encoding/binary"
	"strconv"
)

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
	str := string(c.Id)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (c *CommentInHB) GetVideoId() int64 {
	str := string(c.VideoId)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (c *CommentInHB) GetUserId() int64 {
	str := string(c.UserId)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (c *CommentInHB) GetConversationId() int64 {
	str := string(c.ConversationId)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (c *CommentInHB) GetLastId() int64 {
	str := string(c.LastId)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (c *CommentInHB) GetToUserId() int64 {
	str := string(c.ToUserId)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (c *CommentInHB) GetContent() string {
	return string(c.Content)
}

func (c *CommentInHB) GetTimestamp() string {
	return string(c.Timestamp)
}
