package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/dal/query"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
	"github.com/TremblingV5/DouTok/pkg/utils"
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

	redisMap := map[string]int{
		misc.ComCntCache:      int(misc.GetConfigNum("Redis.ComCntCache.Num")),
		misc.ComTotalCntCache: int(misc.GetConfigNum("Redis.ComTotalCntCache.Num")),
	}
	InitRedis(
		misc.GetConfig("Redis.Dest"),
		misc.GetConfig("Redis.Password"),
		redisMap,
	)

	InitHB(misc.GetConfig("HBase.Host"))

	InitMemoryMap()

	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))
}

func InitRedis(dest string, password string, dbs map[string]int) error {
	redisCaches, _ := redishandle.InitRedis(
		dest, password, dbs,
	)

	RedisClients = redisCaches

	return nil
}

func InitMemoryMap() {
	comCount := safeMap.New()
	comTotalCount := safeMap.New()

	ComCount = comCount
	ComTotalCount = comTotalCount
}

func InitHB(host string) error {
	client := hbaseHandle.InitHB(host)
	HBClient = &client

	return nil
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

	Comment = query.Comment
	CommentCnt = query.CommentCount

	DoComment = Comment.WithContext(context.Background())
	DoCommentCnt = CommentCnt.WithContext(context.Background())

	return nil
}
