package main

import (
	"github.com/TremblingV5/DouTok/applications/comment/handler"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/applications/comment/service"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
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

	svr := commentservice.NewServer(
		new(handler.CommentServiceImpl),
		initHelper.InitRPCServerArgs(misc.Config)...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
