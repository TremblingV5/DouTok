package service

import (
	"github.com/TremblingV5/DouTok/applications/comment/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DoComment query.ICommentDo
var DoCommentCnt query.ICommentCountDo

var Comment = query.CommentStruct
var CommentCnt = query.CommentCountStruct

var HBClient *hbaseHandle.HBaseClient

var RedisClients map[string]*redishandle.RedisClient

// ComCount 在内存中创建一个map用于存储视频的评论数
var ComCount *safeMap.SafeMap
var ComTotalCount *safeMap.SafeMap

// ComContent 达到一定评论数的视频的评论被存储在内存中
var ComContent *safeMap.SafeMap
