package main

import (
	"context"
	"go.uber.org/zap"

	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
)

var (
	logger *zap.Logger
	config = &configStruct.Config{}
)

func init() {
	ctx := context.Background()
	cfg, err := configStruct.Load[*configStruct.Config](ctx, &configStruct.Config{})
	config = cfg

	config.InitViper("DOUTOK_API", "user")
	config.ResolveViperConfig()

	logger = DouTokLogger.InitLogger(config.Logger)
	DouTokContext.DefaultLogger = logger
	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}
}

func main() {
	clients := rpc.New(services.InitRPCClientArgs(constants.USER_SERVER_NAME, config.Etcd))

	options, shutdown := services.InitRPCServerArgs(constants.USER_SERVER_NAME, config.Base, config.Etcd, config.Otel)
	defer shutdown()

	svr := userservice.NewServer(handler.New(clients), options...)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
