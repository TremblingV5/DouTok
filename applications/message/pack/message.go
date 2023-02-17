package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"time"
)

type Message struct {
	Id         int64  `json:"id"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

func NewMessage(req *message.DouyinMessageActionRequest) *Message {
	message := Message{
		Id:         int64(utils.GetSnowFlakeId()),
		FromUserId: req.UserId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
		CreateTime: string(time.Now().Unix()),
	}
	return &message
}
