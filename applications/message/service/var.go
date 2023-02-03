package service

import (
	"github.com/TremblingV5/DouTok/applications/message/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Do query.IMessageDo
var Message = query.MessageStruct

var HBClient *hbaseHandle.HBaseClient
