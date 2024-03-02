package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favorite/handler"
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite/favoriteservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"go.uber.org/zap"
)

//type Config struct {
//	Base   configStruct.Base   `envPrefix:"DOUTOK_FAVORITE_"`
//	Etcd   configStruct.Etcd   `envPrefix:"DOUTOK_FAVORITE_"`
//	Jwt    configStruct.Jwt    `envPrefix:"DOUTOK_FAVORITE_"`
//	Otel   configStruct.Otel   `envPrefix:"DOUTOK_FAVORITE_"`
//	Logger configStruct.Logger `envPrefix:"DOUTOK_FAVORITE_"`
//}

var (
	logger *zap.Logger
	config = &configStruct.Config{}
)

func init() {
	ctx := context.Background()
	cfg, err := configStruct.Load[*configStruct.Config](ctx, &configStruct.Config{})
	config = cfg

	config.InitViper("DOUTOK_API", "favorite")
	config.ResolveViperConfig()

	logger = DouTokLogger.InitLogger(config.Logger)
	DouTokContext.DefaultLogger = logger
	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}
}

func main() {
	options, shutdown := services.InitRPCServerArgs(constants.FAVORITE_SERVER_NAME, config.Base, config.Etcd, config.Otel)
	defer shutdown()

	svr := favoriteservice.NewServer(
		handler.New(rpc.New(services.InitRPCClientArgs(constants.FAVORITE_SERVER_NAME, config.Etcd))),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
