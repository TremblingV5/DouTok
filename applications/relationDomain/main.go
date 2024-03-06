package main

import (
	"github.com/TremblingV5/DouTok/applications/relationDomain/handler"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain/relationdomainservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

var (
	Logger = dlog.InitLog(3)
)

func init() {
	service.Init("relationDomain")
}

func main() {

	options, shutdown := initHelper.InitRPCServerArgs(service.ViperConfig)
	defer shutdown()

	svr := relationdomainservice.NewServer(
		new(handler.RelationDomainServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
