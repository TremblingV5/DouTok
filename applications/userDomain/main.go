package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/userDomain/errs"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"go.uber.org/zap"
	"reflect"

	"github.com/TremblingV5/DouTok/applications/userDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/repository/user"
	"github.com/TremblingV5/DouTok/applications/userDomain/handler"
	"github.com/TremblingV5/DouTok/applications/userDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain/userdomainservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
)

type Config struct {
	Server configStruct.Base
	Etcd   configStruct.Etcd
	MySQL  configStruct.MySQL
	Otel   configStruct.Otel
}

type LoggerConfig struct {
	Logger configStruct.Logger `envPrefix:"DOUTOK_USER_DOMAIN_"`
}

var (
	logger           *zap.Logger
	userDomainConfig Config
	logcfg           LoggerConfig
	viperConfig      *dtviper.Config
)

func init() {
	ctx := context.Background()

	userDomainConfig = Config{}
	logcfg = LoggerConfig{}
	viperConfig = dtviper.ConfigInit("DOUTOK_USER_DOMAIN", "userDomain")
	viperConfig.UnmarshalStructTags(reflect.TypeOf(userDomainConfig), "")
	viperConfig.UnmarshalStruct(&userDomainConfig)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	errs.Init(userDomainConfig.Server)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	DouTokContext.AddLoggerToContext(ctx, logger)

	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}

	logger = DouTokContext.Extract(ctx)
}

func loadFeature() *handler.Handler {
	db, err := userDomainConfig.MySQL.InitDB()
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)

	userRepo := user.New(db)
	userService := service.New(userRepo)
	return handler.New(userService)
}

func main() {

	options, shutdown := services.InitRPCServerArgs(constants.USER_DOMAIN_SERVER_NAME, userDomainConfig.Server, userDomainConfig.Etcd, userDomainConfig.Otel)
	defer shutdown()

	svr := userdomainservice.NewServer(
		loadFeature(),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("failed to run server", zap.Any("error", err))
	}
}
