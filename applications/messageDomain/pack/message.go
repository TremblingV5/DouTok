package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

type Message struct {
	Id         int64  `json:"id"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type HBMessage struct {
	Id         []byte `json:"id"`
	FromUserId []byte `json:"from_user_id"`
	ToUserId   []byte `json:"to_user_id"`
	Content    []byte `json:"content"`
	CreateTime []byte `json:"create_time"`
}

func NewMessage(req *messageDomain.DoutokAddMessageRequest) *entity.Message {
	message := entity.Message{
		Id:         int64(utils.GetSnowFlakeId()),
		FromUserId: req.UserId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
		CreateTime: time.Now().Unix(),
	}
	return &message
}

func HBMsg2Msg(msg *HBMessage) *entity.Message {
	id, err := strconv.ParseInt(string(msg.Id), 10, 64)
	if err != nil {
		klog.Errorf("hbmsg to msg error, err = %s", err)
	}
	fromUserId, err := strconv.ParseInt(string(msg.FromUserId), 10, 64)
	if err != nil {
		klog.Errorf("hbmsg to msg error, err = %s", err)
	}
	toUserId, err := strconv.ParseInt(string(msg.ToUserId), 10, 64)
	if err != nil {
		klog.Errorf("hbmsg to msg error, err = %s", err)
	}
	createTime, err := strconv.ParseInt(string(msg.CreateTime), 10, 64)
	if err != nil {
		klog.Errorf("hbmsg to msg error, err = %s", err)
	}
	return &entity.Message{
		Id:         id,
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		CreateTime: createTime,
		Content:    string(msg.Content),
	}
}
