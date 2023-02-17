package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/dal/query"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
)

func InitDb() error {
	var config configStruct.MySQLConfig
	configurator.InitConfig(
		&config, "mysql.yaml",
	)

	db, err := mysqlIniter.InitDb(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

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

func InitRedis() error {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)

	redisCaches, _ := redishandle.InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)

	RedisClients = redisCaches

	return nil
}

func InitMemoryMap() {
	favCount := safeMap.New()
	favRelationU2V := safeMap.New()
	favRelationV2U := safeMap.New()

	FavCount = favCount
	FavRelationU2V = favRelationU2V
	FavRelationV2U = favRelationV2U
}

func InitKafka() {
	fav_count_kafka_producer := kafka.InitSynProducer([]string{"150.158.237.39:50004"})
	fav_count_kafka_consumer := kafka.InitConsumerGroup([]string{"150.158.237.39:50004"}, misc.FavCountTopicName)

	FavCountKafkaProducer = &fav_count_kafka_producer
	FavCountKafkaConsumer = &fav_count_kafka_consumer
}
