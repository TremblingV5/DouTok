package service

import (
	"github.com/TremblingV5/DouTok/applications/favorite/dal/query"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Dof query.IFavRelationDo
var Dov query.IVideoCountDo
var FavRelation = query.FavRelation
var VideoCount = query.VideoCount
