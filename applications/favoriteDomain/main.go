package main

import (
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/handler"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain/favoritedomainservice"
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

	options, shutdown := services.InitRPCServerArgs(constants.FAVORITE_DOMAIN_SERVER_NAME, service.DomainConfig.BaseConfig)
	defer shutdown()

	svr := favoritedomainservice.NewServer(
		new(handler.FavoriteDomainServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
