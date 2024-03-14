package main

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/TremblingV5/DouTok/applications/publish/handler"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/publish/publishservice"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/services"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_PUBLISH_"`
}

var (
	Logger = dlog.InitLog(3)
	config = &Config{}
)

func init() {
	_, err := configurator.Load(config, "DOUTOK_PUBLISH", "publish")
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}

	rpc.Init()
}

func main() {
	options, shutdown := services.InitRPCServerArgs(constants.PUBLISH_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
