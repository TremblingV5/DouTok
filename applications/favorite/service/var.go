package service

import (
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/favorite/dal/query"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DoFavorite query.IFavoriteDo
var DoFavoriteCnt query.IFavoriteCountDo

var Favorite = query.FavoriteStruct
var FavoriteCnt = query.FavoriteCountStruct

var RedisClients map[string]*redishandle.RedisClient

// 在内存中创建一个map用于存储视频的喜欢数
var FavCount *safeMap.SafeMap
var FavTotalCount *safeMap.SafeMap

// 达到一定喜欢数的视频的喜欢关系被存储在内存中
var FavRelationU2V *safeMap.SafeMap
var FavRelationV2U *safeMap.SafeMap

var FavCountKafkaProducer sarama.SyncProducer
var FavCountKafkaConsumer sarama.ConsumerGroup
