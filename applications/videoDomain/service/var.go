package service

import (
	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/query"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/minioHandle"
	"github.com/TremblingV5/DouTok/pkg/ossHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	minio "github.com/minio/minio-go/v6"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Do query.IVideoDo
var Video = query.VideoStruct

var VideoCount = query.VideoCountStruct
var VideoCountDo query.IVideoCountDo

var HBClient *hbaseHandle.HBaseClient

var OSSClient = &ossHandle.OssClient{
	Client: &oss.Client{},
}
var OssCfg *configStruct.OssConfig

var MinioClient = &minioHandle.MinioClient{
	Client: &minio.Client{},
}

var RedisClients map[string]*redishandle.RedisClient
