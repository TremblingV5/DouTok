package main

import (
	"github.com/TremblingV5/DouTok/applications/message/handler"
	"github.com/TremblingV5/DouTok/applications/message/rpc"
	"github.com/TremblingV5/DouTok/applications/message/service"
	"github.com/TremblingV5/DouTok/kitex_gen/message/messageservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	service.InitViper()
	service.InitHB()
	service.InitRedisClient()
	service.InitSyncProducer()
	service.InitConsumerGroup()
	service.InitId()
	rpc.Init()
}

func main() {
	Init()
	defer func() {
		_ = service.SyncProducer.Close()
	}()

	// 启动 kafka 消费者协程，消费 IM 消息
	go service.ConsumeMsg()

	var logger = dlog.InitLog(3)
	defer logger.Sync()

	klog.SetLogger(logger)

	options, shutdown := initHelper.InitRPCServerArgs(service.ViperConfig)
	defer shutdown()

	svr := messageservice.NewServer(
		new(handler.MessageServiceImpl),
		options...,
	)

	if err := svr.Run(); err != nil {
		klog.Fatal(err)
	}
}
