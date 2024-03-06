package service

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"reflect"
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
	Server configStruct.Base
	MySQL  configStruct.MySQL
	Etcd   configStruct.Etcd
	Redis  configStruct.Redis
	HBase  configStruct.HBase
}

var (
	RedisClient    *redis.Client
	SyncProducer   sarama.SyncProducer
	ViperConfig    *dtviper.Config
	ConsumerGroup  sarama.ConsumerGroup
	ConcurrentMap  *safeMap.SafeMap
	mu             *sync.Mutex
	RelationConfig Config
)

func Init(server string) {
	InitViper(server)
	InitRedisClient()
	InitSyncProducer()
	InitConsumerGroup()
	InitId()
	InitDB()
	InitSafeMap()
	InitMutex()

	go Flush()
}

func InitMutex() {
	mu = &sync.Mutex{}
}

func InitViper(server string) {
	ViperConfig = dtviper.ConfigInit("DOUTOK_RELATION", server)
	ViperConfig.UnmarshalStructTags(reflect.TypeOf(RelationConfig), "")
	ViperConfig.UnmarshalStruct(&RelationConfig)
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
	RedisClient = RelationConfig.Redis.InitRedisClient(configStruct.DEFAULT_DATABASE)
}

func InitId() {
	node := ViperConfig.Viper.GetInt64("Snowflake.Node")
	utils.InitSnowFlake(node)
}

func InitDB() {
	db, err := RelationConfig.MySQL.InitDB()
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)
}

func InitSafeMap() {
	ConcurrentMap = safeMap.New()
}
