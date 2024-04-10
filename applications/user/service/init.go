package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/pkg/utils"

	"github.com/TremblingV5/DouTok/applications/user/dal/query"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
)

func Init() {
	misc.InitViperConfig()

	InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.Database"),
	)

	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))
}

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
