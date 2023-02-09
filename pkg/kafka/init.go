package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

func InitSynProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 等待所有follower都回复ack，确保Kafka不会丢消息
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner // 对Key进行Hash，同样的Key每次都落到一个分区，这样消息是有序的

	cfg := dtviper.ConfigInit("KAFKA", "kafka")

	// 使用同步producer，异步模式下有更高的性能，但是处理更复杂，这里建议先从简单的入手
	producer, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%d", cfg.Viper.GetString("Host"), cfg.Viper.GetInt("Port"))}, config)
	if err != nil {
		panic(err.Error())
	}
	return producer
}

func InitConsumerGroup() sarama.ConsumerGroup {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V2_8_0_0 // specify appropriate version
	consumerConfig.Consumer.Return.Errors = false
	//consumerConfig.Consumer.Offsets.AutoCommit.Enable = true      // 禁用自动提交，改为手动
	//consumerConfig.Consumer.Offsets.AutoCommit.Interval = time.Second * 1 // 测试3秒自动提交
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	cfg := dtviper.ConfigInit("KAFKA", "kafka")

	cGroup, err := sarama.NewConsumerGroup([]string{fmt.Sprintf("%s:%d", cfg.Viper.GetString("Host"), cfg.Viper.GetString("Port"))}, cfg.Viper.GetString("GroupId"), consumerConfig)
	if err != nil {
		panic(err)
	}
	return cGroup
}
