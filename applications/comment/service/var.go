package service

import (
	"github.com/TremblingV5/DouTok/applications/comment/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DoComment query.ICommentDo
var DoCommentCnt query.ICommentCountDo

var Comment = query.CommentStruct
var CommentCnt = query.CommentCountStruct

var HBClient *hbaseHandle.HBaseClient
