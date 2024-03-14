package main

import (
	"github.com/TremblingV5/DouTok/applications/messageDomain/handler"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain/messagedomainservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/TremblingV5/DouTok/applications/messageDomain/service"
	"github.com/TremblingV5/DouTok/pkg/dlog"
)

var (
	Logger = dlog.InitLog(3)
)

func init() {
	service.InitViper()
	service.InitHB()
	service.InitRedisClient()
	service.InitSyncProducer()
	service.InitConsumerGroup()
	service.InitId()
}

func main() {
	options, shutdown := services.InitRPCServerArgs(constants.MESSAGE_DOMAIN_SERVER_NAME, service.DomainConfig.BaseConfig)
	defer shutdown()

	svr := messagedomainservice.NewServer(
		new(handler.MessageDomainServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
