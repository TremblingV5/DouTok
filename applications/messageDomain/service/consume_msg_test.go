package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/pkg/misc"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"log"
	"math"
	"testing"
	"time"
)

func TestConsumeMsg(t *testing.T) {
	Init()

	// 消费4个消息
	go ConsumeMsg()

	time.Sleep(3 * time.Second)

	// 查看 hbase 数据
	sessionId := utils.GenerateSessionId(10001000, 10002000)
	start := fmt.Sprintf("%s%d", sessionId, 0)
	end := fmt.Sprintf("%s%d", sessionId, math.MaxInt)
	res, err := HBClient.ScanRange(ViperConfig.Viper.GetString("Hbase.Table"), start, end)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res {
		hbMsg := pack.HBMessage{}
		err := misc.Map2Struct4HB(v, &hbMsg)
		if err != nil {
			fmt.Println(err)
		}
		assert.DeepEqual(t, "10001000", string(hbMsg.FromUserId))
		assert.DeepEqual(t, "10002000", string(hbMsg.ToUserId))
		fmt.Printf("id = %s, content = %s, createTime = %s\n", string(hbMsg.Id), string(hbMsg.Content), string(hbMsg.CreateTime))
	}

	// 查看 redis 数据
	val, err := RedisClient.HGet(context.Background(), sessionId, "content").Result()
	if err != nil {
		log.Fatal(err)
	}
	assert.DeepEqual(t, "test msg", val)
}
