package main

import (
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/misc"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
)

func main() {
	misc.InitViperConfig()

	if err := service.InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.Database"),
	); err != nil {
		panic(err)
	}

	if err := service.DB.AutoMigrate(&model.Favorite{}, &model.FavoriteCount{}); err != nil {
		panic(err)
	}
}
