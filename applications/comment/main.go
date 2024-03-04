package main

import (
	"github.com/TremblingV5/DouTok/applications/comment/handler"
	"github.com/TremblingV5/DouTok/applications/comment/initialize"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

//type Config struct {
//	Base   configStruct.Base   `envPrefix:"DOUTOK_COMMENT_"`
//	Etcd   configStruct.Etcd   `envPrefix:"DOUTOK_COMMENT_"`
//	Jwt    configStruct.Jwt    `envPrefix:"DOUTOK_COMMENT_"`
//	Otel   configStruct.Otel   `envPrefix:"DOUTOK_COMMENT_"`
//	Logger configStruct.Logger `envPrefix:"DOUTOK_COMMENT_"`
//}

//var (
//	logger *zap.Logger
//	config = &configStruct.Config{}
//)
//
//func init() {
//	ctx := context.Background()
//	cfg, err := configStruct.Load[*configStruct.Config](ctx, &configStruct.Config{})
//	config = cfg
//
//	config.InitViper("DOUTOK_API", "comment")
//	config.ResolveViperConfig()
//	logger = DouTokLogger.InitLogger(config.Logger)
//	DouTokContext.DefaultLogger = logger
//	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
//	if err != nil {
//		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
//	}
//}

func main() {
	initialize.InitComment()
	options, shutdown := services.InitRPCServerArgs(constants.COMMENT_SERVER_NAME, initialize.CommentConfig.Server, initialize.CommentConfig.Etcd, initialize.CommentConfig.Otel)
	defer shutdown()

	svr := commentservice.NewServer(
		handler.New(rpc.New(services.InitRPCClientArgs(constants.COMMENT_SERVER_NAME, initialize.CommentConfig.Etcd))),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
