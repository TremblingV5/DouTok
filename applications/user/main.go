package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/services"
	"go.uber.org/zap"
)

type Config struct {
	Base      configStruct.Base      `envPrefix:"DOUTOK_USER_"`
	Etcd      configStruct.Etcd      `envPrefix:"DOUTOK_USER_"`
	Jwt       configStruct.Jwt       `envPrefix:"DOUTOK_USER_"`
	MySQL     configStruct.MySQL     `envPrefix:"DOUTOK_USER_"`
	Snowflake configStruct.Snowflake `envPrefix:"DOUTOK_USER_"`
	HBase     configStruct.HBase     `envPrefix:"DOUTOK_USER_"`
	Redis     configStruct.Redis     `envPrefix:"DOUTOK_USER_"`
	Otel      configStruct.Otel      `envPrefix:"DOUTOK_USER_"`
	Logger    configStruct.Logger    `envPrefix:"DOUTOK_USER_"`
}

var (
	logger *zap.Logger
	config = &Config{}
)

func init() {
	ctx := context.Background()
	cfg, err := configStruct.Load[*Config](ctx, &Config{})
	config = cfg
	logger = DouTokLogger.InitLogger(config.Logger)
	DouTokContext.DefaultLogger = logger
	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}
}

func main() {
	clients := rpc.New(config.Base.GetName(), services.InitRPCClientArgs(config.Base, config.Etcd))

	options, shutdown := services.InitRPCServerArgs(config.Base, config.Etcd, config.Otel)
	defer shutdown()

	svr := userservice.NewServer(handler.New(clients), options...)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
