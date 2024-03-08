package main

import (
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
)

func main() {

	service.Init()

	if err := service.DB.AutoMigrate(&model.Favorite{}, &model.FavoriteCount{}); err != nil {
		panic(err)
	}
}
