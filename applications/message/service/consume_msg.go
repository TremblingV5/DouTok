package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/pkg/misc"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

type msgConsumerGroup struct{}

func (m msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (m msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (m msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		message := pack.Message{}
		json.Unmarshal(msg.Value, &message)
		mp, err := misc.Struct2Map(message)
		if err != nil {
			return err
		}
		// Struct2Map 有bug，会转float64
		klog.Infof("messages to map, msg content = %s from %d to %d\n", mp["content"], int64(mp["from_user_id"].(float64)), int64(mp["to_user_id"].(float64)))
		sessionId := utils.GenerateSessionId(message.FromUserId, message.ToUserId)

		// 更新 redis 的最新消息
		err = RedisClient.HMSet(context.Background(), sessionId, mp).Err()
		if err != nil {
			return err
		}

		content, err := RedisClient.HGet(context.Background(), sessionId, "content").Result()
		if err != nil {
			klog.Errorf("get friend list message error, err = %s", err)
		}
		klog.Infof("content = %s\n", content)

		// 将消息存入 hbase
		// 生成 rowkey
		rowKey := fmt.Sprintf("%s%d", sessionId, message.CreateTime)

		println("consume msg form kafka, generate rowKey = ", rowKey)

		// 构造 hbase 一条数据
		hbData := map[string]map[string][]byte{
			"data": {
				"id":           []byte(fmt.Sprintf("%d", message.Id)),
				"from_user_id": []byte(fmt.Sprintf("%d", message.FromUserId)),
				"to_user_id":   []byte(fmt.Sprintf("%d", message.ToUserId)),
				"content":      []byte(message.Content),
				"create_time":  []byte(fmt.Sprintf("%d", message.CreateTime)),
			},
		}

		err = HBClient.Put(ViperConfig.Viper.GetString("Hbase.Table"), rowKey, hbData)
		if err != nil {
			return err
		}

		// TODO 1、解决消息重复消费 2、分布式事务一致性（redis 与 hbase 同时成功或失败、重试策略）

		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")

	}
	return nil
}

var consumerGroup msgConsumerGroup

func ConsumeMsg() {

	for {
		err := ConsumerGroup.Consume(context.Background(), ViperConfig.Viper.GetStringSlice("Kafka.Topics"), consumerGroup)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	_ = ConsumerGroup.Close()
}
