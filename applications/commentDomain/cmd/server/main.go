package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/handler"
	"github.com/TremblingV5/DouTok/applications/commentDomain/redis/commentTotalCountRedis"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain/commentdomainservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"go.uber.org/zap"
	"reflect"
)

type Config struct {
	Server    configStruct.Base
	Etcd      configStruct.Etcd
	Otel      configStruct.Otel
	MySQL     configStruct.MySQL
	Snowflake configStruct.Snowflake
	HBase     configStruct.HBase
	Redis     configStruct.Redis
}

type LoggerConfig struct {
	Logger configStruct.Logger `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
}

var (
	logger              *zap.Logger
	commentDomainConfig Config
	logcfg              LoggerConfig
	viperConfig         *dtviper.Config
	handle              *handler.CommentDomainHandler
)

func init() {
	ctx := context.Background()
	commentDomainConfig = Config{}
	logcfg = LoggerConfig{}
	viperConfig = dtviper.ConfigInit("DOUTOK_COMMENT_DOMAIN", "commentDomain")
	viperConfig.UnmarshalStructTags(reflect.TypeOf(commentDomainConfig), "")
	viperConfig.UnmarshalStruct(&commentDomainConfig)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}

	logger = DouTokContext.Extract(ctx)

	db, err := commentDomainConfig.MySQL.InitDB()
	if err != nil {
		panic(err)
	}

	hb := hbaseHandle.HBaseClient{
		Client: commentDomainConfig.HBase.InitHB(),
	}

	redisClient := redishandle.RedisClient{
		Client: commentDomainConfig.Redis.InitRedisClient(1),
	}
	commentTotalCountRedisClient := commentTotalCountRedis.NewClient(&redisClient)
	commentDomainService := service.NewCommentDomainService(
		db, &hb, commentTotalCountRedisClient, commentDomainConfig.Snowflake.Node,
	)

	handle = handler.NewCommentDomainHandler(commentDomainService)
}

func main() {

	options, shutdown := initHelper.InitRPCServerArgs(viperConfig)
	defer shutdown()

	svr := commentdomainservice.NewServer(
		handle, options...,
	)
	if err := svr.Run(); err != nil {
		logger.Fatal("start server defeat", zap.Error(err))
	}
}
