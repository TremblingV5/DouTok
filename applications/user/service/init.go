package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/dal/query"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
)

func InitDb() error {
	var config configStruct.MySQLConfig
	configurator.InitConfig(
		&config, "mysql.yaml",
	)

	db, err := mysqlIniter.InitDb(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
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
