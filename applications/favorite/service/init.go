package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/dal/query"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
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
