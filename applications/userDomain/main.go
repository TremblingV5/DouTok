package main

import (
	"context"
	"go.uber.org/zap"

	"github.com/TremblingV5/DouTok/applications/userDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/repository/user"
	"github.com/TremblingV5/DouTok/applications/userDomain/errs"
	"github.com/TremblingV5/DouTok/applications/userDomain/handler"
	"github.com/TremblingV5/DouTok/applications/userDomain/service"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain/userdomainservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_USER_DOMAIN_"`
	Jwt                     configStruct.Jwt    `envPrefix:"DOUTOK_USER_DOMAIN_"`
	MySQL                   configStruct.MySQL  `envPrefix:"DOUTOK_USER_DOMAIN_"`
	Logger                  configStruct.Logger `envPrefix:"DOUTOK_USER_DOMAIN_"`
}

var (
	logger *zap.Logger
	config = &Config{}
)

func init() {
	ctx := context.Background()

	_, err := configurator.Load(config, "DOUTOK_USER_DOMAIN", "userDomain")
	errs.Init(config.Base)
	logger = DouTokLogger.InitLogger(config.Logger)
	DouTokContext.DefaultLogger = logger
	DouTokContext.AddLoggerToContext(ctx, logger)

	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}

	logger = DouTokContext.Extract(ctx)
}

func loadFeature() *handler.Handler {
	db, err := config.MySQL.InitDB()
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)

	userRepo := user.New(db)
	userService := service.New(userRepo)
	return handler.New(userService)
}

func main() {
	options, shutdown := services.InitRPCServerArgs(constants.USER_DOMAIN_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := userdomainservice.NewServer(
		loadFeature(),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("failed to run server", zap.Any("error", err))
	}
}
