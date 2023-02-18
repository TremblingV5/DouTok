package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/dal/query"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"github.com/TremblingV5/DouTok/pkg/safeMap"
)

func InitMemoryMap() {
	comCount := safeMap.New()
	comContent := safeMap.New()

	ComCount = comCount
	ComContent = comContent
}

func InitHB(host string) error {
	client := hbaseHandle.InitHB(host)
	HBClient = &client

	return nil
}

func InitDb(username string, password string, host string, port string, database string) error {
	db, err := mysqlIniter.InitDb(
		username,
		password,
		host,
		port,
		database,
	)

	if err != nil {
		return err
	}

	DB = db

	query.SetDefault(DB)

	Comment = query.Comment
	CommentCnt = query.CommentCount

	DoComment = Comment.WithContext(context.Background())
	DoCommentCnt = CommentCnt.WithContext(context.Background())

	return nil
}
