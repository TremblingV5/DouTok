package main

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/comment/handler"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
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
	Logger configStruct.Logger `envPrefix:"DOUTOK_COMMENT_"`
}

var (
	logger        *zap.Logger
	commentConfig Config
	logcfg        LoggerConfig
	ViperConfig   *dtviper.Config
)

func init() {
	ctx := context.Background()
	commentConfig = Config{}
	logcfg = LoggerConfig{}
	ViperConfig = dtviper.ConfigInit("DOUTOK_COMMENT", "comment")
	ViperConfig.UnmarshalStructTags(reflect.TypeOf(commentConfig), "")
	ViperConfig.UnmarshalStruct(&commentConfig)
	fmt.Println(commentConfig.Etcd)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}
}

func main() {

	options, shutdown := services.InitRPCServerArgs(constants.COMMENT_SERVER_NAME, commentConfig.Server, commentConfig.Etcd, commentConfig.Otel)
	defer shutdown()

	svr := commentservice.NewServer(
		handler.New(rpc.New(services.InitRPCClientArgs(constants.COMMENT_SERVER_NAME, commentConfig.Etcd))),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
