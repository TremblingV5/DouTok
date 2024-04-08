package main

import (
	"github.com/TremblingV5/DouTok/applications/publish/dal/model"
	"github.com/TremblingV5/DouTok/applications/publish/misc"
	"github.com/TremblingV5/DouTok/applications/publish/service"
)

func main() {
	misc.InitViperConfig()

	service.InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.Database"),
	)

	service.DB.AutoMigrate(&model.Video{})
}
