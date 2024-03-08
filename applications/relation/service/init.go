package service

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/query"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_RELATION_"`
	Redis                   configStruct.Redis `envPrefix:"DOUTOK_RELATION_"`
	MySQL                   configStruct.MySQL `envPrefix:"DOUTOK_RELATION_"`
}

var (
	RedisClient   *redis.Client
	SyncProducer  sarama.SyncProducer
	ViperConfig   *dtviper.Config
	ConsumerGroup sarama.ConsumerGroup
	ConcurrentMap *safeMap.SafeMap
	mu            *sync.Mutex
	DomainConfig  = &Config{}
)

func InitMutex() {
	mu = &sync.Mutex{}
}

func InitViper() {
	v, err := configurator.Load(DomainConfig, "DOUTOK_RELATION", "relation")
	ViperConfig = v
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", DomainConfig))
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
	RedisClient = DomainConfig.Redis.InitRedisClient(configStruct.DEFAULT_DATABASE)
}

func InitId() {
	node := ViperConfig.Viper.GetInt64("Snowflake.Node")
	utils.InitSnowFlake(node)
}

func InitDB() {
	db, err := DomainConfig.MySQL.InitDB()
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)
}

func InitSafeMap() {
	ConcurrentMap = safeMap.New()
}
