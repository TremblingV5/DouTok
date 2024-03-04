package main

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/commentDomain/cmd/server/initialize"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain/commentdomainservice"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

//type Config struct {
//	Base      configStruct.Base      `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	Etcd      configStruct.Etcd      `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	Jwt       configStruct.Jwt       `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	MySQL     configStruct.MySQL     `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	Snowflake configStruct.Snowflake `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	HBase     configStruct.HBase     `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	Redis     configStruct.Redis     `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	Otel      configStruct.Otel      `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//	Logger    configStruct.Logger    `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
//}
//
//var (
//	logger *zap.Logger
//	handle = &handler.CommentDomainHandler{}
//	config = &Config{}
//)
//
//func init() {
//	ctx := context.Background()
//
//	cfg, err := configStruct.Load[*Config](ctx, &Config{})
//	config = cfg
//	logger = DouTokLogger.InitLogger(config.Logger)
//	DouTokContext.DefaultLogger = logger
//	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
//
//	if err != nil {
//		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
//	}
//
//	logger = DouTokContext.Extract(ctx)
//
//	db, err := mysqlIniter.InitDb(
//		config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, strconv.Itoa(config.MySQL.Port), config.MySQL.Database,
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	hb := hbaseHandle.InitHB(config.HBase.Host)
//
//	redisClient := redishandle.NewRedisClient(config.Redis.Dsn, config.Redis.Password, 1)
//	commentTotalCountRedisClient := commentTotalCountRedis.NewClient(redisClient)
//	commentDomainService := service.NewCommentDomainService(
//		db, &hb, commentTotalCountRedisClient, config.Snowflake.Node,
//	)
//
//	handle = handler.NewCommentDomainHandler(commentDomainService)
//}

func main() {
	initialize.InitCommentDomain()
	fmt.Println(initialize.CommentDomainConfig.Redis)
	options, shutdown := initHelper.InitRPCServerArgsV2(
		initialize.CommentDomainConfig.Server, initialize.CommentDomainConfig.Etcd, initialize.CommentDomainConfig.Otel,
	)
	defer shutdown()

	svr := commentdomainservice.NewServer(
		initialize.Handle, options...,
	)
	if err := svr.Run(); err != nil {
		logger.Fatal("start server defeat", zap.Error(err))
	}
}
