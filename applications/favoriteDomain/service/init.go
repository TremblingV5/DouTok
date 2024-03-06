package service

import (
	"context"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"reflect"

	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/misc"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

type Config struct {
	MySQL configStruct.MySQL
	Redis configStruct.Redis
}

var favoriteDomainConfig Config

func Init() {

	v := dtviper.ConfigInit(misc.ViperConfigEnvPrefix, misc.ViperConfigEnvFilename)
	v.UnmarshalStructTags(reflect.TypeOf(favoriteDomainConfig), "")
	v.UnmarshalStruct(&favoriteDomainConfig)
	FavoriteConfig = v

	if err := InitDb(); err != nil {
		panic(err)
	}

	redisMap := map[string]int{
		misc.FavCache:         v.Viper.GetInt(misc.ConfigIndex_RedisFavCacheDbNum),
		misc.FavCntCache:      v.Viper.GetInt(misc.ConfigIndex_RedisFavCntCacheDbNum),
		misc.FavTotalCntCache: v.Viper.GetInt("Redis.FavTotalCountCache.Num"),
	}

	InitRedis(redisMap)

	InitMemoryMap()

	kafkaBrokers := []string{
		v.Viper.GetString("Kafka.Broker"),
	}
	InitKafka(kafkaBrokers)

	utils.InitSnowFlake(
		v.Viper.GetInt64(misc.ConfigIndex_SnowFlake),
	)
}

func InitDb() error {
	db, err := favoriteDomainConfig.MySQL.InitDB()
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
			Client: favoriteDomainConfig.Redis.InitRedisClient(v),
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
