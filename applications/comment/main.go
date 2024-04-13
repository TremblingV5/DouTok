package main

import (
	"github.com/TremblingV5/DouTok/applications/comment/api/comment_api"
	"github.com/TremblingV5/DouTok/applications/comment/infra/misc"
	"github.com/TremblingV5/DouTok/applications/comment/infra/query"
	"github.com/TremblingV5/DouTok/applications/comment/infra/repository/comment_hb_repo"
	"github.com/TremblingV5/DouTok/applications/comment/infra/rpc"
	"github.com/TremblingV5/DouTok/applications/comment/services/comment_service"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
	"github.com/TremblingV5/DouTok/pkg/cache"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/box/dbtx"
)

var (
	Logger = dlog.InitLog(3)
)

func initDb(
	username, password, host, port, database string,
) {
	db, err := mysqlIniter.InitDb(
		username,
		password,
		host,
		port,
		database,
	)

	if err != nil {
		panic(err)
	}

	query.SetDefault(db)

	dbtx.Init(func() dbtx.TX {
		return query.Q.Begin()
	})
}

func initRedis(
	dest, password string,
	dbs map[string]int,
) map[string]*redishandle.RedisClient {
	redisClients, err := redishandle.InitRedis(
		dest,
		password,
		dbs,
	)

	if err != nil {
		panic(err)
	}

	return redisClients
}

func initHb(host string) hbaseHandle.HBaseClient {
	return hbaseHandle.InitHB(host)
}

func Init() *comment_api.CommentServiceImpl {
	misc.InitViperConfig()

	initDb(
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
	redisClients := initRedis(
		misc.GetConfig("Redis.Dest"),
		misc.GetConfig("Redis.Password"),
		redisMap,
	)

	hbaseClient := initHb(misc.GetConfig("HBase.Host"))

	service := comment_service.New(
		cache.NewCountMapCache(),
		cache.NewCountMapCache(),
		comment_hb_repo.New(&hbaseClient),
		redisClients[misc.ComCntCache],
		redisClients[misc.ComTotalCntCache],
	)

	go service.UpdateComCountMap()
	go service.UpdateComTotalCntMap()

	handle := comment_api.New(service, rpc.New(initHelper.InitRPCClientArgs(misc.Config)))
	return handle
}

func main() {
	Init()

	options, shutdown := initHelper.InitRPCServerArgs(misc.Config)
	defer shutdown()

	svr := commentservice.NewServer(
		Init(),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
