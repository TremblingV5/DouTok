package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"reflect"
)

type Config struct {
	Server configStruct.Base
	Etcd   configStruct.Etcd
	MySQL  configStruct.MySQL
	Redis  configStruct.Redis
	HBase  configStruct.HBase
}

var (
	ViperConfig *dtviper.Config
	VideoConfig Config
)

func Init() {
	ViperConfig = dtviper.ConfigInit(misc.ViperConfigEnvPrefix, misc.ViperConfigEnvFilename)
	ViperConfig.UnmarshalStructTags(reflect.TypeOf(VideoConfig), "")
	ViperConfig.UnmarshalStruct(&VideoConfig)

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
		misc.SendBox:    ViperConfig.Viper.GetInt("Redis.SendBox.Num"),
		misc.MarkedTime: ViperConfig.Viper.GetInt("Redis.MarkedTime.Num"),
	}

	InitRedis(redisMap)
}

func InitRedis(dbs map[string]int) {
	RedisClients = make(map[string]*redishandle.RedisClient)
	for k, v := range dbs {
		RedisClients[k] = &redishandle.RedisClient{
			Client: VideoConfig.Redis.InitRedisClient(v),
		}
	}
}

func InitDb() error {
	db, err := VideoConfig.MySQL.InitDB()

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
		Client: VideoConfig.HBase.InitHB(),
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
