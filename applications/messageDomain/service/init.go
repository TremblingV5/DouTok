package service

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_MESSAGEDOMAIN_"`
	Redis                   configStruct.Redis `envPrefix:"DOUTOK_MESSAGEDOMAIN_"`
	HBase                   configStruct.HBase `envPrefix:"DOUTOK_MESSAGEDOMAIN_"`
}

var (
	HBClient      *hbaseHandle.HBaseClient
	RedisClient   *redis.Client
	SyncProducer  sarama.SyncProducer
	ViperConfig   *dtviper.Config
	ConsumerGroup sarama.ConsumerGroup
	DomainConfig  = &Config{}
)

func InitViper() {
	v, err := configurator.Load(DomainConfig, "DOUTOK_MESSAGEDOMAIN", "messageDomain")
	ViperConfig = v
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", DomainConfig))
	}
}

func InitHB() {
	HBClient = &hbaseHandle.HBaseClient{Client: *DomainConfig.HBase.InitHB()}
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
	RedisClient = DomainConfig.Redis.InitRedisClient(configStruct.DEFAULT_DATABASE)
}

func InitId() {
	node := ViperConfig.Viper.GetInt64("Snowflake.Node")
	utils.InitSnowFlake(node)
}
