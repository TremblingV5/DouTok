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
}

var (
	ViperConfig  *dtviper.Config
	DomainConfig Config
	Logger       *LogBuilder.Logger
)

func Init() {

	Logger = LogBuilder.New("./tmp/videoDomain.log", 1024*1024, 3, 10)

	v, err := configurator.Load(&DomainConfig, "DOUTOK_VIDEODOMAIN", "videoDomain")
	ViperConfig = v
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", DomainConfig))
	}

	if err := InitDb(); err != nil {
		panic(err)
	}

	InitHB()

	if err := InitMinio(
		ViperConfig.Viper.GetString("Minio.Endpoint"),
		ViperConfig.Viper.GetString("Minio.Key"),
		ViperConfig.Viper.GetString("Minio.Secret"),
		ViperConfig.Viper.GetString("Minio.Bucket"),
	); err != nil {
		panic(nil)
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

func InitOSS(endpoint string, key string, secret string, bucketName string) error {
	if err := OSSClient.Init(
		endpoint, key, secret, bucketName,
	); err != nil {
		return err
	}

	config := configStruct.OssConfig{
		Endpoint:   endpoint,
		Key:        key,
		Secret:     secret,
		BucketName: bucketName,
		//Callback:   callback,
	}

	OssCfg = &config

	return nil
}

func InitMinio(endpoint string, key string, secret string, bucketName string) error {
	MinioClient.Init(
		endpoint, key, secret, bucketName,
	)

	config := configStruct.OssConfig{
		Endpoint:   endpoint,
		Key:        key,
		Secret:     secret,
		BucketName: bucketName,
		//Callback:   callback,
	}

	OssCfg = &config

	return nil
}
