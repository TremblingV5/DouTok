package main

import (
	"github.com/TremblingV5/DouTok/applications/relationDomain/handler"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain/relationdomainservice"
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

func loadFeature() *handler.Handler {
	// TODO xban 加载 repo 等其他对象
	relationService := service.New()
	return handler.New(relationService)
}

func main() {

	options, shutdown := services.InitRPCServerArgs(constants.RELATION_DOMAIN_SERVER_NAME, service.DomainConfig.BaseConfig)
	defer shutdown()

	svr := relationdomainservice.NewServer(
		loadFeature(),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
