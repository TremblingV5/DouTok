package main

import (
	"github.com/TremblingV5/DouTok/applications/relation/dal/model"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
)

func main() {

	cfg := dtviper.ConfigInit("", "relation")

	db, err := mysqlIniter.InitDb(
		cfg.Viper.GetString("MySQL.Username"),
		cfg.Viper.GetString("MySQL.Password"),
		cfg.Viper.GetString("MySQL.Host"),
		cfg.Viper.GetString("MySQL.Port"),
		cfg.Viper.GetString("MySQL.Database"),
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Relation{}, &model.FollowCount{}, &model.FollowerCount{})
}
