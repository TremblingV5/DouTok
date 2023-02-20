package service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/relation/dal/query"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	RedisClient   *redis.Client
	SyncProducer  sarama.SyncProducer
	ViperConfig   dtviper.Config
	ConsumerGroup sarama.ConsumerGroup
	ConcurrentMap *safeMap.SafeMap
	mu            *sync.Mutex
)

func InitMutex() {
	mu = &sync.Mutex{}
}

func InitViper() {
	ViperConfig = dtviper.ConfigInit("DOUTOK_RELATION", "relation")
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

func InitDB() {
	username := ViperConfig.Viper.GetString("MySQL.Username")
	password := ViperConfig.Viper.GetString("MySQL.Password")
	host := ViperConfig.Viper.GetString("MySQL.Host")
	port := ViperConfig.Viper.GetString("MySQL.Port")
	database := ViperConfig.Viper.GetString("MySQL.Database")
	db, err := mysqlIniter.InitDb(username, password, host, port, database)
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)
}

func InitSafeMap() {
	ConcurrentMap = safeMap.New()
}
