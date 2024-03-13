package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/TremblingV5/DouTok/applications/favorite/handler"
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite/favoriteservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/dtx"
	"github.com/TremblingV5/DouTok/pkg/services"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_FAVORITE_"`
	Jwt                     configStruct.Jwt    `envPrefix:"DOUTOK_FAVORITE_"`
	Logger                  configStruct.Logger `envPrefix:"DOUTOK_FAVORITE_"`
}

var (
	logger *zap.Logger
	config = &Config{}
)

func init() {
	ctx := context.Background()

	favoriteConfig = Config{}
	_, err := configurator.Load(config, "DOUTOK_FAVORITE", "favorite")
	logger = DouTokLogger.InitLogger(config.Logger)
	dtx.DefaultLogger = logger
	dtx.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}
}

func main() {
	options, shutdown := services.InitRPCServerArgs(constants.FAVORITE_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := favoriteservice.NewServer(
		handler.New(rpc.New(services.InitRPCClientArgs(constants.FAVORITE_SERVER_NAME, config.Etcd))),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
