package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/misc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"go.uber.org/zap"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_FAVORITEDOMAIN_"`
	Redis                   configStruct.Redis `envPrefix:"DOUTOK_FAVORITEDOMAIN_"`
	MySQL                   configStruct.MySQL `envPrefix:"DOUTOK_FAVORITEDOMAIN_"`
	HBase                   configStruct.HBase `envPrefix:"DOUTOK_FAVORITEDOMAIN_"`
}

var (
	ViperConfig  *dtviper.Config
	DomainConfig = &Config{}
)

func Init() {

	v, err := configurator.Load(DomainConfig, "DOUTOK_FAVORITEDOMAIN", "favoriteDomain")
	ViperConfig = v
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", DomainConfig))
	}

	if err := InitDb(); err != nil {
		panic(err)
	}

	redisMap := map[string]int{
		misc.FavCache:         ViperConfig.Viper.GetInt(misc.ConfigIndex_RedisFavCacheDbNum),
		misc.FavCntCache:      ViperConfig.Viper.GetInt(misc.ConfigIndex_RedisFavCntCacheDbNum),
		misc.FavTotalCntCache: ViperConfig.Viper.GetInt("Redis.FavTotalCountCache.Num"),
	}

	InitRedis(redisMap)

	InitMemoryMap()

	kafkaBrokers := []string{
		ViperConfig.Viper.GetString("Kafka.Broker"),
	}
	InitKafka(kafkaBrokers)

	utils.InitSnowFlake(
		ViperConfig.Viper.GetInt64(misc.ConfigIndex_SnowFlake),
	)
}

func InitDb() error {
	db, err := DomainConfig.MySQL.InitDB()

	if err != nil {
		return err
	}

	DB = db

	query.SetDefault(DB)

	Favorite = query.Favorite
	FavoriteCnt = query.FavoriteCount

	DoFavorite = Favorite.WithContext(context.Background())
	DoFavoriteCnt = FavoriteCnt.WithContext(context.Background())

	return nil
}

func InitRedis(dbs map[string]int) {
	RedisClients = make(map[string]*redishandle.RedisClient)
	for k, v := range dbs {
		RedisClients[k] = &redishandle.RedisClient{
			Client: DomainConfig.Redis.InitRedisClient(v),
		}
	}
}

func InitMemoryMap() {
	favCount := safeMap.New()
	favTotalCount := safeMap.New()
	favRelationU2V := safeMap.New()
	favRelationV2U := safeMap.New()

	FavCount = favCount
	FavTotalCount = favTotalCount
	FavRelationU2V = favRelationU2V
	FavRelationV2U = favRelationV2U
}

func InitKafka(brokers []string) {
	fav_count_kafka_producer := kafka.InitSynProducer(brokers)
	fav_count_kafka_consumer := kafka.InitConsumerGroup(brokers, misc.FavCountGroupName)

	FavCountKafkaProducer = fav_count_kafka_producer
	FavCountKafkaConsumer = fav_count_kafka_consumer
}
