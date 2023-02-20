package service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/go-redis/redis/v8"
)

var (
	HBClient      *hbaseHandle.HBaseClient
	RedisClient   *redis.Client
	SyncProducer  sarama.SyncProducer
	ViperConfig   dtviper.Config
	ConsumerGroup sarama.ConsumerGroup
)

func InitViper() {
	ViperConfig = dtviper.ConfigInit("DOUTOK_MESSAGE", "message")
}

func InitHB() {

	client := hbaseHandle.InitHB(ViperConfig.Viper.GetString("Hbase.Host"))

	HBClient = &client
}

func InitSyncProducer() {
	producer := kafka.InitSynProducer(ViperConfig.Viper.GetStringSlice("Kafka.Brokers"))
	SyncProducer = producer
}

func InitConsumerGroup() {
	cGroup := kafka.InitConsumerGroup(ViperConfig.Viper.GetStringSlice("Kafka.Brokers"), ViperConfig.Viper.GetStringSlice("Kafka.GroupIds")[0])
	ConsumerGroup = cGroup
}

func InitRedisClient() {

	Client, err := redishandle.InitRedisClient(
		fmt.Sprintf("%s:%d", ViperConfig.Viper.GetString("Redis.Host"), ViperConfig.Viper.GetInt("Redis.Port")),
		ViperConfig.Viper.GetString("Redis.Password"),
		ViperConfig.Viper.GetInt("Redis.Databases.Default"),
	)
	if err != nil {
		panic(err)
	}
	RedisClient = Client
}

func InitId() {
	node := ViperConfig.Viper.GetInt64("Snowflake.Node")
	utils.InitSnowFlake(node)
}
