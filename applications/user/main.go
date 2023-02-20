package main

import (
	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/applications/user/service"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
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

	svr := userservice.NewServer(
		new(handler.UserServiceImpl),
		initHelper.InitRPCServerArgs(misc.Config)...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
