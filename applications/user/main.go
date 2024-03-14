package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dtx"
	"github.com/TremblingV5/DouTok/pkg/services"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_USER_"`
	Logger                  configStruct.Logger `envPrefix:"DOUTOK_USER_"`
}

var (
	logger *zap.Logger
	config = &Config{}
)

func init() {
	ctx := context.Background()
	
	_, err := configurator.Load(config, "DOUTOK_USER", "user")
	logger = DouTokLogger.InitLogger(config.Logger)
	dtx.DefaultLogger = logger
	dtx.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}
}

func main() {
	clients := rpc.New(services.InitRPCClientArgs(constants.USER_SERVER_NAME, config.Etcd))

	options, shutdown := services.InitRPCServerArgs(constants.USER_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := userservice.NewServer(handler.New(clients), options...)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
