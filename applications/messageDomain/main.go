package main

import (
	"github.com/TremblingV5/DouTok/applications/messageDomain/handler"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain/messagedomainservice"

	// "github.com/TremblingV5/DouTok/applications/messageDomain/misc"
	"github.com/TremblingV5/DouTok/applications/messageDomain/service"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

var (
	Logger = dlog.InitLog(3)
)

func Init() {
	// misc.InitViperConfig()
	service.InitViper()
	service.InitHB()
	service.InitRedisClient()
	service.InitSyncProducer()
	service.InitConsumerGroup()
	service.InitId()
}

func main() {
	Init()

	options, shutdown := initHelper.InitRPCServerArgs(service.ViperConfig)
	defer shutdown()

	svr := messagedomainservice.NewServer(
		new(handler.MessageDomainServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
