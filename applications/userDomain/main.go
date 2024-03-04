package main

import (
	"github.com/TremblingV5/DouTok/applications/userDomain/initialize"
	"github.com/bytedance/gopkg/util/logger"
	"strconv"

	"go.uber.org/zap"

	"github.com/TremblingV5/DouTok/applications/userDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/repository/user"
	"github.com/TremblingV5/DouTok/applications/userDomain/handler"
	"github.com/TremblingV5/DouTok/applications/userDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain/userdomainservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"github.com/TremblingV5/DouTok/pkg/services"
)

//type Config struct {
//	Base      configStruct.Base      `envPrefix:"DOUTOK_USER_DOMAIN_"`
//	Etcd      configStruct.Etcd      `envPrefix:"DOUTOK_USER_DOMAIN_"`
//	Jwt       configStruct.Jwt       `envPrefix:"DOUTOK_USER_DOMAIN_"`
//	MySQL     configStruct.MySQL     `envPrefix:"DOUTOK_USER_DOMAIN_"`
//	Snowflake configStruct.Snowflake `envPrefix:"DOUTOK_USER_DOMAIN_"`
//	Otel      configStruct.Otel      `envPrefix:"DOUTOK_USER_DOMAIN_"`
//	Logger    configStruct.Logger    `envPrefix:"DOUTOK_USER_DOMAIN_"`
//}

//var (
//	logger *zap.Logger
//	config = &configStruct.Config{}
//)
//
//func init() {
//	ctx := context.Background()
//
//	cfg, err := configStruct.Load[*configStruct.Config](ctx, &configStruct.Config{})
//	config = cfg
//
//	config.InitViper("DOUTOK_API", "userDomain")
//	config.ResolveViperConfig()
//
//	errs.Init(config.Base)
//	logger = DouTokLogger.InitLogger(config.Logger)
//	DouTokContext.DefaultLogger = logger
//	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
//
//	if err != nil {
//		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
//	}
//
//	logger = DouTokContext.Extract(ctx)
//}

func loadFeature() *handler.Handler {
	db, err := mysqlIniter.InitDb(
		initialize.UserDomainConfig.MySQL.Username,
		initialize.UserDomainConfig.MySQL.Password,
		initialize.UserDomainConfig.MySQL.Host,
		strconv.Itoa(initialize.UserDomainConfig.MySQL.Port),
		initialize.UserDomainConfig.MySQL.Database,
	)
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)

	userRepo := user.New(db)
	userService := service.New(userRepo)
	return handler.New(userService)
}

func main() {
	initialize.InitUserDomain()
	options, shutdown := services.InitRPCServerArgs(constants.USER_DOMAIN_SERVER_NAME, initialize.UserDomainConfig.Server, initialize.UserDomainConfig.Etcd, initialize.UserDomainConfig.Otel)
	defer shutdown()

	svr := userdomainservice.NewServer(
		loadFeature(),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("failed to run server", zap.Any("error", err))
	}
}
