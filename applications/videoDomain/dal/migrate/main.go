package main

import (
	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"reflect"
)

type Config struct {
	MySQL configStruct.MySQL
}

var migrateConfig Config

func main() {
	v := dtviper.ConfigInit(misc.ViperConfigEnvPrefix, misc.ViperConfigEnvFilename)
	v.UnmarshalStructTags(reflect.TypeOf(migrateConfig), "")
	v.UnmarshalStruct(&migrateConfig)

	if _, err := migrateConfig.MySQL.InitDB(); err != nil {
		panic(err)
	}

	if err := service.DB.AutoMigrate(&model.Video{}, &model.VideoCount{}); err != nil {
		panic(err)
	}
}
