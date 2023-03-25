package service

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/messageDomain/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *messageDomain.DoutokAddMessageRequest) error {
	// 使用同步producer，将消息存入 kafka
	// 构建消息
	val, err := json.Marshal(pack.NewMessage(req))
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: ViperConfig.Viper.GetStringSlice("Kafka.Topics")[0],
		Value: sarama.StringEncoder(val),
	}
	partition, offset, err := SyncProducer.SendMessage(msg)

	if err == nil {
		klog.Infof("produce success, partition: %d, offset: %d\n", partition, offset)
	} else {
		return err
	}

	return nil
}
