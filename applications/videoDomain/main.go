package main

import (
	"github.com/TremblingV5/DouTok/applications/videoDomain/handler"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain/videodomainservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/services"
)

var (
	Logger = dlog.InitLog(3)
)

func init() {
	service.Init()
}

func main() {

	options, shutdown := services.InitRPCServerArgs(constants.VIDEO_DOMAIN_SERVER_NAME, service.DomainConfig.BaseConfig)
	defer shutdown()

	svr := videodomainservice.NewServer(
		new(handler.VideoDomainServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
