package service

import (
	"github.com/TremblingV5/DouTok/applications/feed/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Do query.IVideoDo
var Video = query.VideoStruct

var RedisClients map[string]*redishandle.RedisClient
var HBClient *hbaseHandle.HBaseClient
