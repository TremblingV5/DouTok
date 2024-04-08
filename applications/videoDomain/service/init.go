package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/query"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/LogBuilder"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_VIDEODOMAIN_"`
	MySQL                   configStruct.MySQL `envPrefix:"DOUTOK_VIDEODOMAIN_"`
	Redis                   configStruct.Redis `envPrefix:"DOUTOK_VIDEODOMAIN_"`
	HBase                   configStruct.HBase `envPrefix:"DOUTOK_VIDEODOMAIN_"`
	MinIO                   configStruct.MinIO `envPrefix:"DOUTOK_VIDEODOMAIN_"`
}

var (
	ViperConfig  *dtviper.Config
	DomainConfig Config
	Logger       *LogBuilder.Logger
)

func Init() {

	Logger = LogBuilder.New("./tmp/videoDomain.log", 1024*1024, 3, 10)

	// TODO 参数和行为不太一致，这个函数做了两件事：初始化DomainConfig，但是返回了ViperConfig
	v, err := configurator.Load(&DomainConfig, "DOUTOK_VIDEODOMAIN", "videoDomain")
	ViperConfig = v
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", DomainConfig))
	}

	if err := InitDb(); err != nil {
		panic(err)
	}

	InitHB()

	if err := InitMinio(); err != nil {
		panic(err.Error())
	}

	utils.InitSnowFlake(ViperConfig.Viper.GetInt64("Snowflake.Node"))
	redisMap := map[string]int{
		"SendBox":    ViperConfig.Viper.GetInt("Redis.SendBox.Num"),
		"MarkedTime": ViperConfig.Viper.GetInt("Redis.MarkedTime.Num"),
	}

	InitRedis(redisMap)
}

func InitRedis(dbs map[string]int) {
	RedisClients = make(map[string]*redishandle.RedisClient)
	for k, v := range dbs {
		RedisClients[k] = &redishandle.RedisClient{
			Client: DomainConfig.Redis.InitRedisClient(v),
		}
	}
}

func InitDb() error {
	db, err := DomainConfig.MySQL.InitDB()

	if err != nil {
		return err
	}

	DB = db

	query.SetDefault(DB)
	Video = query.Video
	Do = Video.WithContext(context.Background())

	VideoCount = query.VideoCount
	VideoCountDo = VideoCount.WithContext(context.Background())

	return nil
}

func InitHB() {
	HBClient = &hbaseHandle.HBaseClient{
		Client: *DomainConfig.HBase.InitHB(),
	}
}

func InitMinio() error {
	client, err := DomainConfig.MinIO.InitIO()
	if err != nil {
		return err
	}
	MinioClient.Client = client
	MinioClient.Bucket = DomainConfig.MinIO.Bucket

	return nil
}
