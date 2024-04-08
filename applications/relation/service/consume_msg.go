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
		userId := fmt.Sprintf("%d", relation.UserId)
		toUserId := fmt.Sprintf("%d", relation.ToUserId)
		actionType := fmt.Sprintf("%d", relation.ActionType)

		fmt.Printf("userId = %s, toUserId = %s, actionType = %s\n", userId, toUserId, actionType)

		// 更新关注表 cache
		err := WriteFollowToCache(userId, toUserId, actionType)
		if err != nil {
			return err
		}
		// 更新粉丝表 cache
		err = WriteFollowerToCache(toUserId, userId, actionType)
		if err != nil {
			return err
		}
		// 更新关注表 db
		err = WriteFollowToDB(&relation)
		if err != nil {
			klog.Errorf("write follow to db error, err = %s", err)
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
		time.Sleep(time.Second * 1)
		klog.Infof("start iter\n")
		ConcurrentMap.Iter(iter)
		// 需要清空
		ConcurrentMap.Clean()
	}
}

func iter(key string, v interface{}) {
	if v == nil {
		return
	}
	fmt.Printf("key = %s, v = %v\n", key, v)
	value := v.(int64)
	pair := strings.Split(key, "-")
	userId, err := strconv.ParseInt(pair[1], 10, 64)
	klog.Infof("userId = %d\n", userId)
	if err != nil {
		klog.Errorf("strconv.ParseInt error, err = %s", err)
	}
	if value != 0 {
		if pair[0][6] == '_' {
			klog.Infof("关注更新 %d\n", value)
			// follow_
			// 更新关注数 db
			UpdateFollowCountFromDB(userId, value)
			// 删除关注数 cache
			err = DeleteFollowCountCache(pair[1])
		} else {
			klog.Infof("粉丝更新 %d\n", value)
			// follower_
			// 更新粉丝数 db
			UpdateFollowerCountFromDB(userId, value)
			// 删除粉丝数 cache
			err = DeleteFollowerCountCache(pair[1])
		}
	}
}
