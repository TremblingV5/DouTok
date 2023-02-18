package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/dal/query"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
)

func InitDb(username string, password string, host string, port string, database string) error {
	db, err := mysqlIniter.InitDb(
		username, password, host, port, database,
	)

	if err != nil {
		return err
	}

	DB = db

	query.SetDefault(DB)

	User = query.User
	Do = User.WithContext(context.Background())

	return nil
}
