package main

import (
	"github.com/TremblingV5/DouTok/applications/publish/handler"
	"github.com/TremblingV5/DouTok/applications/publish/misc"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/applications/publish/service"
	"github.com/TremblingV5/DouTok/kitex_gen/publish/publishservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

var (
	Logger = dlog.InitLog(3)
)

func Init() {
	misc.InitViperConfig()
	service.Init()
	rpc.InitPRCClient()
}

func main() {
	Init()

	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		initHelper.InitRPCServerArgs(misc.Config)...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
