package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"time"
)

type Message struct {
	Id          int64  `json:"id"`
	FromUserId  int64  `json:"from_user_id"`
	ToUserId    int64  `json:"to_user_id"`
	Content     string `json:"content"`
	ContentType int32  `json:"content_type"`
	CreateTime  string `json:"create_time"`
}

func NewMessage(msg *message.DouyinRelationActionRequest) *Message {
	message := Message{
		Id:          int64(utils.GetSnowFlakeId()),
		FromUserId:  msg.UserId,
		ToUserId:    msg.ToUserId,
		Content:     msg.Content,
		ContentType: msg.ActionType,
		CreateTime:  string(time.Now().Unix()),
	}
	return &message
}
