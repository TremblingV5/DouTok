package service

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"reflect"

	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Server configStruct.Base
	Etcd   configStruct.Etcd
	Redis  configStruct.Redis
	HBase  configStruct.HBase
}

var (
	HBClient      *hbaseHandle.HBaseClient
	RedisClient   *redis.Client
	SyncProducer  sarama.SyncProducer
	ViperConfig   *dtviper.Config
	ConsumerGroup sarama.ConsumerGroup
	MessageConfig Config
)

func InitViper() {
	ViperConfig = dtviper.ConfigInit("DOUTOK_MESSAGE", "message")
	ViperConfig.UnmarshalStructTags(reflect.TypeOf(MessageConfig), "")
	ViperConfig.UnmarshalStruct(&MessageConfig)
}

func InitHB() {
	HBClient = &hbaseHandle.HBaseClient{
		Client: MessageConfig.HBase.InitHB(),
	}
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
	RedisClient = MessageConfig.Redis.InitRedisClient(configStruct.DEFAULT_DATABASE)
}

func InitId() {
	node := ViperConfig.Viper.GetInt64("Snowflake.Node")
	utils.InitSnowFlake(node)
}
