package service

import (
	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/minioHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	minio "github.com/minio/minio-go/v6"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Do query.IVideoDo
var Video = query.VideoStruct

var VideoCount = query.VideoCountStruct
var VideoCountDo query.IVideoCountDo

var HBClient *hbaseHandle.HBaseClient

var MinioClient = &minioHandle.MinioClient{
	Client: &minio.Client{},
}

var RedisClients map[string]*redishandle.RedisClient
