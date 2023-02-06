package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// 使用同步producer，异步模式下有更高的性能，但是处理更复杂，这里建议先从简单的入手
	producer := kafka.InitSynProducer()
	defer func() {
		_ = producer.Close()
	}()

	msgCount := 4
	// 模拟4个消息
	for i := 0; i < msgCount; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		msg := &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.StringEncoder("hello+" + strconv.Itoa(rand.Int())),
			Key:   sarama.StringEncoder("BBB"),
		}

		t1 := time.Now().Nanosecond()
		partition, offset, err := producer.SendMessage(msg)
		t2 := time.Now().Nanosecond()

		if err == nil {
			fmt.Println("produce success, partition:", partition, ",offset:", offset, ",cost:", (t2-t1)/(1000*1000), " ms")
		} else {
			fmt.Println(err.Error())
		}
	}
}
