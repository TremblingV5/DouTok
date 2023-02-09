package service

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *message.DouyinRelationActionRequest) error {
	// 使用同步producer，将消息存入 kafka
	defer func() {
		_ = SyncProducer.Close()
	}()
	// 构建消息
	val, err := json.Marshal(pack.NewMessage(req))
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: "message",
		Value: sarama.StringEncoder(val),
		Key:   sarama.StringEncoder("key"),
	}
	partition, offset, err := SyncProducer.SendMessage(msg)

	if err == nil {
		klog.Infof("produce success, partition:", partition, ",offset:", offset)
	} else {
		return err
	}

	return nil
}
