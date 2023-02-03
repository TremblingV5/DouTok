package service

import (
	"github.com/TremblingV5/DouTok/applications/publish/dal/query"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/ossHandle"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Do query.IVideoDo
var Video = query.VideoStruct

var HBClient *hbaseHandle.HBaseClient

var OSSClient = &ossHandle.OssClient{
	Client: &oss.Client{},
}
var OssCfg *configStruct.OssConfig
