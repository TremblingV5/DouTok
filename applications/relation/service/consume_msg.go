package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"strings"
	"time"
)

type msgConsumerGroup struct{}

func (m msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (m msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (m msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		klog.Infof("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		relation := pack.Relation{}
		json.Unmarshal(msg.Value, &relation)
		// 更新关注表 cache
		err := WriteFollowToCache(string(relation.UserId), string(relation.ToUserId), string(relation.ActionType))
		if err != nil {
			return err
		}
		// 更新粉丝表 cache
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
		// 更新粉丝数（db）
		err = UpdateFollowerCountFromDB(relation.ToUserId, op)
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
		err := ConsumerGroup.Consume(context.Background(), ViperConfig.Viper.GetStringSlice("Kafka.Topics"), consumerGroup)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	_ = ConsumerGroup.Close()
}

/**
1. 从 SafeMap 中获取关注数和粉丝数，累加入 MySQL
2. 删除 Redis 中对应的缓存 follow/follower 数量
*/
func Flush() {
	for {
		time.Sleep(time.Second * 3)
		ConcurrentMap.Iter(iter)
		// 需要清空
		ConcurrentMap.Clean()
	}
}

func iter(key string, v interface{}) {
	value := v.(int64)
	pair := strings.Split(key, "-")
	user_id, err := strconv.ParseInt(pair[1], 10, 64)
	if err != nil {
		klog.Errorf("strconv.ParseInt error, err = %s", err)
	}
	if value != 0 {
		if pair[0][6] == '_' {
			// follow_
			// 更新关注数 db
			UpdateFollowCountFromDB(user_id, value)
			// 删除关注数 cache
			err = DeleteFollowCountCache(pair[1])
		} else {
			// follower_
			// 更新粉丝数 db
			UpdateFollowerCountFromDB(user_id, value)
			// 删除粉丝数 cache
			err = DeleteFollowerCountCache(pair[1])
		}
	}
}
