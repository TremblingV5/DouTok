package main

import (
	"github.com/TremblingV5/DouTok/applications/feed/handler"
	"github.com/TremblingV5/DouTok/applications/feed/misc"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/feed/feedservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

var (
	Logger     = dlog.InitLog(3)
	feedConfig *dtviper.Config
)

func init() {
	feedConfig = dtviper.ConfigInit(misc.ViperConfigEnvPrefix, misc.ViperConfigEnvFilename)

	rpc.Init()
}

func main() {

	options, shutdown := initHelper.InitRPCServerArgs(feedConfig)
	defer shutdown()

	svr := feedservice.NewServer(
		new(handler.FeedServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
