package main

import (
	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/applications/user/initialize"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func main() {
	initialize.InitUser()
	clients := rpc.New(services.InitRPCClientArgs(constants.USER_SERVER_NAME, initialize.UserConfig.Etcd))

	options, shutdown := services.InitRPCServerArgs(constants.USER_SERVER_NAME, initialize.UserConfig.Server, initialize.UserConfig.Etcd, initialize.UserConfig.Otel)
	defer shutdown()

	svr := userservice.NewServer(handler.New(clients), options...)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
