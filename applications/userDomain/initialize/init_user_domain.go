package initialize

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/userDomain/errs"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"go.uber.org/zap"
	"reflect"
)

type Config struct {
	Server configStruct.Base
	Etcd   configStruct.Etcd
	MySQL  configStruct.MySQL
	Otel   configStruct.Otel
	//Jwt       configStruct.Jwt
	//Snowflake configStruct.Snowflake
}

type LoggerConfig struct {
	Logger configStruct.Logger `envPrefix:"DOUTOK_USER_DOMAIN_"`
}

var (
	logger           *zap.Logger
	UserDomainConfig Config
	logcfg           LoggerConfig
	ViperConfig      *dtviper.Config
)

func InitUserDomain() {
	ctx := context.Background()

	UserDomainConfig = Config{}
	logcfg = LoggerConfig{}
	ViperConfig = dtviper.ConfigInit("DOUTOK_USER_DOMAIN", "userDomain", reflect.TypeOf(UserDomainConfig))
	ViperConfig.UnmarshalStruct(&UserDomainConfig)
	fmt.Println(UserDomainConfig.Server.Name)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	errs.Init(UserDomainConfig.Server)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	ctx = DouTokContext.AddLoggerToContext(ctx, logger)

	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}

	logger = DouTokContext.Extract(ctx)
}
