package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

func Init() {
	misc.InitViperConfig()

	InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.Database"),
	)
	InitHB(
		misc.GetConfig("HBase.Host"),
	)
	InitMinio(
		misc.GetConfig("Minio.Endpoint"),
		misc.GetConfig("Minio.Key"),
		misc.GetConfig("Minio.Secret"),
		misc.GetConfig("Minio.Bucket"),
	)
	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))
	redisMap := map[string]int{
		misc.SendBox:    int(misc.GetConfigNum("Redis.SendBox.Num")),
		misc.MarkedTime: int(misc.GetConfigNum("Redis.MarkedTime.Num")),
	}
	InitRedis(
		misc.GetConfig("Redis.Dest"),
		misc.GetConfig("Redis.Password"),
		redisMap,
	)
}

func InitRedis(dest string, password string, dbs map[string]int) error {
	redisCaches, _ := redishandle.InitRedis(
		dest, password, dbs,
	)

	RedisClients = redisCaches

	return nil
}

func InitDb(username string, password string, host string, port string, database string) error {
	db, err := mysqlIniter.InitDb(
		username, password, host, port, database,
	)

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

func InitHB(host string) error {
	client := hbaseHandle.InitHB(host)
	HBClient = &client

	return nil
}

func InitOSS(endpoint string, key string, secret string, bucketName string) error {
	OSSClient.Init(
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
