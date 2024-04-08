package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/handler"
	"github.com/TremblingV5/DouTok/applications/comment/redis/commentTotalCountRedis"
	"github.com/TremblingV5/DouTok/applications/comment/service"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"go.uber.org/zap"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_COMMENT_"`
	Jwt                     configStruct.Jwt       `envPrefix:"DOUTOK_COMMENT_"`
	Logger                  configStruct.Logger    `envPrefix:"DOUTOK_COMMENT_"`
	HBase                   configStruct.HBase     `envPrefix:"DOUTOK_COMMENT_"`
	MySQL                   configStruct.MySQL     `envPrefix:"DOUTOK_COMMENT_"`
	Redis                   configStruct.Redis     `envPrefix:"DOUTOK_COMMENT_"`
	Snowflake               configStruct.Snowflake `envPrefix:"DOUTOK_COMMENT_"`
}

var (
	logger *zap.Logger
	config = &Config{}
	handle = &handler.Handler{}
)

func init() {
	ctx := context.Background()
	_, err := configurator.Load(config, "DOUTOK_COMMENT", "comment")
	logger = DouTokLogger.InitLogger(config.Logger)
	DouTokContext.DefaultLogger = logger
	DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}

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
	commentDomainService := service.NewcommentService(
		db, &hb, commentTotalCountRedisClient, config.Snowflake.Node,
	)

	utils.InitSnowFlake(config.Snowflake.Node)

	handle = handler.New(commentDomainService)
}

func main() {
	options, shutdown := services.InitRPCServerArgs(constants.COMMENT_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := commentservice.NewServer(
		handle,
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
