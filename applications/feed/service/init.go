package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/feed/misc"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/utils"

	"github.com/TremblingV5/DouTok/applications/feed/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
)

func Init() {
	misc.InitViperConfig()

	InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.Database"),
	)

	InitHB(
		misc.GetConfig("HBase.Host"),
	)

	redisMap := map[string]int{
		constants.FeedSendBox: int(misc.GetConfigNum("Redis.SendBox.Num")),
		constants.TimeCache:   int(misc.GetConfigNum("Redis.MarkdedTime.Num")),
	}
	InitRedis(
		misc.GetConfig("Redis.Dest"),
		misc.GetConfig("Redis.Password"),
		redisMap,
	)
	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))
}

func InitDb(username string, password string, host string, port string, database string) error {
	db, err := mysqlIniter.InitDb(
		username,
		password,
		host,
		port,
		database,
	)

	if err != nil {
		return err
	}

	DB = db

	query.SetDefault(DB)
	Video = query.Video
	Do = Video.WithContext(context.Background())

	return nil
}

func InitHB(host string) error {
	client := hbaseHandle.InitHB(host)
	HBClient = &client

	return nil
}

func InitRedis(dest string, password string, dbs map[string]int) error {
	redisCaches, _ := redishandle.InitRedis(
		dest, password, dbs,
	)

	RedisClients = redisCaches

	return nil
}
