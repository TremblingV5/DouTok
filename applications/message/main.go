package main

import (
	"go.uber.org/zap"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/TremblingV5/DouTok/applications/message/rpc"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/TremblingV5/DouTok/applications/message/handler"
	"github.com/TremblingV5/DouTok/applications/message/service"
	"github.com/TremblingV5/DouTok/kitex_gen/message/messageservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
)

func init() {
	service.InitViper()
	service.InitHB()
	service.InitRedisClient()
	service.InitSyncProducer()
	service.InitConsumerGroup()
	service.InitId()
	rpc.Init()
}

func main() {
	defer func() {
		_ = service.SyncProducer.Close()
	}()

	// 启动 kafka 消费者协程，消费 IM 消息
	go service.ConsumeMsg()

	var logger = dlog.InitLog(3)
	defer logger.Sync()

	klog.SetLogger(logger)

	options, shutdown := services.InitRPCServerArgs(constants.MESSAGE_SERVER_NAME, service.DomainConfig.BaseConfig)
	defer shutdown()

	svr := messageservice.NewServer(
		new(handler.MessageServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
