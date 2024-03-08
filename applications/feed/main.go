package main

import (
	"github.com/TremblingV5/DouTok/applications/feed/handler"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/feed/feedservice"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_FEED_"`
}

var (
	Logger = dlog.InitLog(3)
	config = &Config{}
)

func init() {

	_, err := configurator.Load(config, "DOUTOK_FEED_", "feed")
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}
	rpc.Init()
}

func main() {

	options, shutdown := services.InitRPCServerArgs(constants.FEED_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := feedservice.NewServer(
		new(handler.FeedServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
