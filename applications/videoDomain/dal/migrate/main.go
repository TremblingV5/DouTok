package main

import (
	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
)

func main() {
	service.Init()

	if err := service.DB.AutoMigrate(&model.Video{}, &model.VideoCount{}); err != nil {
		panic(err)
	}
}
