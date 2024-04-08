package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/publish/dal/query"
	"github.com/TremblingV5/DouTok/applications/publish/misc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
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
	InitOSS(
		misc.GetConfig("OSS.Endpoint"),
		misc.GetConfig("OSS.Key"),
		misc.GetConfig("OSS.Secret"),
		misc.GetConfig("OSS.Bucket"),
	)
	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))
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
