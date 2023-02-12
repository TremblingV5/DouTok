package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/dal/query"
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
	FavRelation = query.FavRelation
	VideoCount = query.VideoCount
	Dof = FavRelation.WithContext(context.Background())
	Dov = VideoCount.WithContext(context.Background())

	return nil
}
