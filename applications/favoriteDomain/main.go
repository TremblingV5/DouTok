package main

import (
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/handler"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain/favoritedomainservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

var (
	Logger = dlog.InitLog(3)
)

func init() {
	service.Init()
}

func main() {

	options, shutdown := initHelper.InitRPCServerArgs(service.FavoriteConfig)
	defer shutdown()

	svr := favoritedomainservice.NewServer(
		new(handler.FavoriteDomainServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
