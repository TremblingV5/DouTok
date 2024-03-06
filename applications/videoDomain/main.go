package main

import (
	"github.com/TremblingV5/DouTok/applications/videoDomain/handler"
	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain/videodomainservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

var (
	Logger = dlog.InitLog(3)
)

func init() {
	misc.InitLogger()
	service.Init()
}

func main() {

	options, shutdown := initHelper.InitRPCServerArgs(service.ViperConfig)
	defer shutdown()

	svr := videodomainservice.NewServer(
		new(handler.VideoDomainServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
