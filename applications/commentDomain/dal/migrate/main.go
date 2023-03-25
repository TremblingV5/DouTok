package main

import (
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
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

	service.DB.AutoMigrate(&model.Comment{}, &model.CommentCount{})
}
