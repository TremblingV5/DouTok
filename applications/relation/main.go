package main

import (
	"github.com/TremblingV5/DouTok/applications/relation/handler"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/applications/relation/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/cloudwego/kitex/pkg/klog"
)

func init() {
	service.InitViper()
	service.InitRedisClient()
	service.InitSyncProducer()
	service.InitConsumerGroup()
	service.InitId()
	service.InitDB()
	service.InitSafeMap()
	service.InitMutex()
}

func main() {

	defer func() {
		_ = service.SyncProducer.Close()
	}()

	// 启动 kafka 消费者协程，消费点赞消息
	go service.ConsumeMsg()

	var logger = dlog.InitLog(3)
	defer logger.Sync()

	klog.SetLogger(logger)

	clients := rpc.New(services.InitRPCClientArgs(constants.USER_SERVER_NAME, service.DomainConfig.Etcd))

	options, shutdown := services.InitRPCServerArgs(constants.RELATION_SERVER_NAME, service.DomainConfig.BaseConfig)
	defer shutdown()

	svr := relationservice.NewServer(
		handler.New(clients),
		options...,
	)

	if err := svr.Run(); err != nil {
		klog.Fatalf("stopped with error:", err)
	}
}
