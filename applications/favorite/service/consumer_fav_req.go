package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
)

type favCountConsumerGroup struct{}

func (m favCountConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (m favCountConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (m favCountConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		req := FavReqInKafka{}
		json.Unmarshal(msg.Value, &req)

		err := CreateFavoriteInRDB(req.UserId, req.VideoId, req.Op)

		if err != nil {
			// TODO: 写日志
			return err
		}

		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

var group favCountConsumerGroup

func Consumer4UpdateCount() {
	for {
		err := FavCountKafkaConsumer.Consume(context.Background(), []string{misc.FavCountTopicName}, group)

		if err != nil {
			break
		}

	}

	_ = FavCountKafkaConsumer.Close()
}
