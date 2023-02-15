package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
)

type msgConsumerGroup struct{}

func (m msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (m msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (m msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		relation := pack.Relation{}
		json.Unmarshal(msg.Value, &relation)
		// 更新关注表 cache
		err := WriteFollowToCache(string(relation.UserId), string(relation.ToUserId), string(relation.ActionType))
		if err != nil {
			return err
		}
		err = WriteFollowerToCache(string(relation.ToUserId), string(relation.UserId), string(relation.ActionType))
		if err != nil {
			return err
		}
		// 更新关注表 db
		err = WriteFollowToDB(&relation)
		if err != nil {
			return err
		}
		op := int64(0)
		if relation.ActionType == 1 {
			op = 1
		} else {
			op = -1
		}
		// 更新关注数（db）
		err = UpdateFollowCountFromDB(relation.UserId, op)
		if err != nil {
			return err
		}
		// 删除关注数缓存
		err = DeleteFollowCountCache(string(relation.UserId))
		if err != nil {
			return err
		}
		// 更新粉丝数（db）
		err = UpdateFollowerCountFromDB(relation.ToUserId, op)
		if err != nil {
			return err
		}
		// 删除粉丝数缓存
		err = DeleteFollowerCountCache(string(relation.UserId))
		if err != nil {
			return err
		}
		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

var consumerGroup msgConsumerGroup

func ConsumeMsg() {

	for {
		err := ConsumerGroup.Consume(context.Background(), []string{"test"}, consumerGroup)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	_ = ConsumerGroup.Close()
}
