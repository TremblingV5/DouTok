package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/pkg/kafka"
)

type msgConsumerGroup struct{}

func (m msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (m msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (m msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		// 其他数据落库操作

		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

var consumerGroup msgConsumerGroup

func main() {

	cGroup := kafka.InitConsumerGroup()

	for {
		err := cGroup.Consume(context.Background(), []string{"test"}, consumerGroup)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	_ = cGroup.Close()
}
