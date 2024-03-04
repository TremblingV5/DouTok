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
	Logger configStruct.Logger `envPrefix:"DOUTOK_FAVORITE_"`
}

var (
	logger         *zap.Logger
	favoriteConfig Config
	logcfg         LoggerConfig
	viperConfig    *dtviper.Config
)

func init() {
	ctx := context.Background()
	favoriteConfig = Config{}
	logcfg = LoggerConfig{}
	viperConfig = dtviper.ConfigInit("DOUTOK_FAVORITE", "favorite")
	viperConfig.UnmarshalStructTags(reflect.TypeOf(favoriteConfig), "")
	viperConfig.UnmarshalStruct(&favoriteConfig)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}
}
func main() {

	options, shutdown := services.InitRPCServerArgs(constants.FAVORITE_SERVER_NAME, favoriteConfig.Server, favoriteConfig.Etcd, favoriteConfig.Otel)
	defer shutdown()

	svr := favoriteservice.NewServer(
		handler.New(rpc.New(services.InitRPCClientArgs(constants.FAVORITE_SERVER_NAME, favoriteConfig.Etcd))),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
