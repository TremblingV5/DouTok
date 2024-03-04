package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/services"
	"go.uber.org/zap"
	"reflect"
)

type Config struct {
	Server configStruct.Base
	Etcd   configStruct.Etcd
	Otel   configStruct.Otel
}

type LoggerConfig struct {
	Logger configStruct.Logger `envPrefix:"DOUTOK_USER_"`
}

var (
	logger      *zap.Logger
	userConfig  Config
	logcfg      LoggerConfig
	viperConfig *dtviper.Config
)

func init() {
	ctx := context.Background()
	userConfig = Config{}
	logcfg = LoggerConfig{}
	viperConfig = dtviper.ConfigInit("DOUTOK_USER", "user")
	viperConfig.UnmarshalStructTags(reflect.TypeOf(userConfig), "")
	viperConfig.UnmarshalStruct(&userConfig)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}
}

func main() {
	clients := rpc.New(services.InitRPCClientArgs(constants.USER_SERVER_NAME, userConfig.Etcd))

	options, shutdown := services.InitRPCServerArgs(constants.USER_SERVER_NAME, userConfig.Server, userConfig.Etcd, userConfig.Otel)
	defer shutdown()

	svr := userservice.NewServer(handler.New(clients), options...)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
