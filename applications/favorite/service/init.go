package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/dal/query"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/pkg/kafka"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
)

func InitDb(username string, password string, host string, port string, database string) error {
	db, err := mysqlIniter.InitDb(
		username, password, host, port, database,
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

func InitRedis(dest string, password string, dbs map[string]int) error {
	redisCaches, _ := redishandle.InitRedis(
		dest, password, dbs,
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

func InitKafka(brokers []string) {
	fav_count_kafka_producer := kafka.InitSynProducer(brokers)
	fav_count_kafka_consumer := kafka.InitConsumerGroup(brokers, misc.FavCountTopicName)

	FavCountKafkaProducer = fav_count_kafka_producer
	FavCountKafkaConsumer = fav_count_kafka_consumer
}
