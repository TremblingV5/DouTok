package initialize

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/handler"
	"github.com/TremblingV5/DouTok/applications/commentDomain/redis/commentTotalCountRedis"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"go.uber.org/zap"
	"reflect"
	"strconv"
)

type Config struct {
	Server configStruct.Base
	Etcd   configStruct.Etcd
	Otel   configStruct.Otel
	//Jwt       configStruct.Jwt
	MySQL     configStruct.MySQL
	Snowflake configStruct.Snowflake
	HBase     configStruct.HBase
	Redis     configStruct.Redis
}

type LoggerConfig struct {
	Logger configStruct.Logger `envPrefix:"DOUTOK_COMMENT_DOMAIN_"`
}

var (
	logger              *zap.Logger
	CommentDomainConfig Config
	logcfg              LoggerConfig
	ViperConfig         *dtviper.Config
	Handle              *handler.CommentDomainHandler
)

func InitCommentDomain() {
	ctx := context.Background()
	CommentDomainConfig = Config{}
	logcfg = LoggerConfig{}
	ViperConfig = dtviper.ConfigInit("DOUTOK_COMMENT_DOMAIN", "commentDomain", reflect.TypeOf(CommentDomainConfig))
	ViperConfig.UnmarshalStruct(&CommentDomainConfig)

	logcfg, err := configStruct.Load[*LoggerConfig](ctx, &logcfg)

	logger = DouTokLogger.InitLogger(logcfg.Logger)
	DouTokContext.DefaultLogger = logger
	ctx = DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", logcfg))
	}

	logger = DouTokContext.Extract(ctx)

	db, err := mysqlIniter.InitDb(
		CommentDomainConfig.MySQL.Username, CommentDomainConfig.MySQL.Password, CommentDomainConfig.MySQL.Host, strconv.Itoa(CommentDomainConfig.MySQL.Port), CommentDomainConfig.MySQL.Database,
	)
	if err != nil {
		panic(err)
	}

	hb := hbaseHandle.InitHB(CommentDomainConfig.HBase.Host)

	redisClient := redishandle.NewRedisClient(CommentDomainConfig.Redis.Dsn, CommentDomainConfig.Redis.Password, 1)
	commentTotalCountRedisClient := commentTotalCountRedis.NewClient(redisClient)
	commentDomainService := service.NewCommentDomainService(
		db, &hb, commentTotalCountRedisClient, CommentDomainConfig.Snowflake.Node,
	)

	Handle = handler.NewCommentDomainHandler(commentDomainService)
}
