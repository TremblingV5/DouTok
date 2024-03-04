package initialize

import (
	"context"
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
	Otel   configStruct.Otel
	//Jwt       configStruct.Jwt
	//MySQL     configStruct.MySQL
	//Snowflake configStruct.Snowflake
	//HBase     configStruct.HBase
	//Redis     configStruct.Redis
}

type LoggerConfig struct {
	Logger configStruct.Logger `envPrefix:"DOUTOK_COMMENT_"`
}

var (
	logger        *zap.Logger
	CommentConfig Config
	logcfg        LoggerConfig
	ViperConfig   *dtviper.Config
)

func InitComment() {
	ctx := context.Background()
	CommentConfig = Config{}
	logcfg = LoggerConfig{}
	ViperConfig = dtviper.ConfigInit("DOUTOK_COMMENT", "comment", reflect.TypeOf(CommentConfig))
	ViperConfig.UnmarshalStruct(&CommentConfig)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}
}
