package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/TremblingV5/DouTok/applications/commentDomain/handler"
	"github.com/TremblingV5/DouTok/applications/commentDomain/redis/commentTotalCountRedis"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain/commentdomainservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/dtx"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/services"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`

	MySQL     configStruct.MySQL     `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
	Snowflake configStruct.Snowflake `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
	HBase     configStruct.HBase     `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
	Redis     configStruct.Redis     `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`

	Logger configStruct.Logger `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
}

var (
	logger *zap.Logger
	handle = &handler.CommentDomainHandler{}
	config = &Config{}
)

func init() {
	ctx := context.Background()

	_, err := configurator.Load(config, "DOUTOK_COMMENT_DOMAIN", "commentDomain")

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	dtx.DefaultLogger = logger
	ctx = dtx.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}

	logger = dtx.Extract(ctx)

	db, err := config.MySQL.InitDB()
	if err != nil {
		panic(err)
	}

	hb := hbaseHandle.HBaseClient{
		Client: *config.HBase.InitHB(),
	}
	redisClient := redishandle.RedisClient{
		Client: config.Redis.InitRedisClient(1),
	}
	commentTotalCountRedisClient := commentTotalCountRedis.NewClient(&redisClient)
	commentDomainService := service.NewCommentDomainService(
		db, &hb, commentTotalCountRedisClient, config.Snowflake.Node,
	)

	handle = handler.NewCommentDomainHandler(commentDomainService)
}

func main() {

	options, shutdown := services.InitRPCServerArgs(constants.COMMENT_DOMAIN_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := commentdomainservice.NewServer(
		handle, options...,
	)
	if err := svr.Run(); err != nil {
		logger.Fatal("start server defeat", zap.Error(err))
	}
}
