package service

import (
	"github.com/TremblingV5/DouTok/applications/favorite/dal/query"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DoFavorite query.IFavoriteDo
var DoFavoriteCnt query.IFavoriteCountDo

var Favorite = query.FavoriteStruct
var FavoriteCnt = query.FavoriteCountStruct

var RedisClients map[string]*redishandle.RedisClient
