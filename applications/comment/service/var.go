package service

import (
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RdClient *redishandle.RedisClient

const RdDefault = "CommentRdDefault"
