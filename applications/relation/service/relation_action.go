package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// 关注 || 取关动作
func (s *RelationActionService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	// 在 SafeMap 中更新局部关注数和粉丝数
	followKey := fmt.Sprintf("%s%d", constants.FollowCount, req.UserId)
	followerKey := fmt.Sprintf("%s%d", constants.FollowerCount, req.ToUserId)
	follow, ok := ConcurrentMap.Get(followKey)
	if !ok {
		klog.Infof("get follow count from concurrentMap false")
	}
	follower, ok := ConcurrentMap.Get(followerKey)
	if !ok {
		klog.Infof("get follow count from concurrentMap false")
	}
	op := int64(0)
	if req.ActionType == 1 {
		op = 1
	} else {
		op = -1
	}
	// TODO 如果关注或者取关对应的增加 safemap 值，前提是需要验证重复性操作
	mu.Lock()
	if follow == nil {
		klog.Infof("set follow %s, %d\n", followKey, op)
		ConcurrentMap.Set(followKey, op)
	} else {
		klog.Infof("set follow %s, %d\n", followKey, follow.(int64)+op)
		ConcurrentMap.Set(followKey, follow.(int64)+op)
	}
	if follower == nil {
		klog.Infof("set follower %s, %d\n", followerKey, op)
		ConcurrentMap.Set(followerKey, op)
	} else {
		klog.Infof("set follower %s, %d\n", followerKey, follower.(int64)+op)
		ConcurrentMap.Set(followerKey, follower.(int64)+op)
	}
	follow, ok = ConcurrentMap.Get(followKey)
	if !ok {
		klog.Errorf("concurrentMap get false")
	}
	follower, ok = ConcurrentMap.Get(followerKey)
	if !ok {
		klog.Errorf("concurrentMap get false")
	}
	klog.Infof("%s follow = %d\n", followKey, follow.(int64))
	klog.Infof("%s follower = %d\n", followerKey, follower.(int64))
	mu.Unlock()
	// 使用同步producer，将消息存入 kafka
	// 构建消息
	val, err := json.Marshal(pack.NewRelation(req))
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
