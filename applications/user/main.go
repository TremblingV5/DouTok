package main

import (
	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

const (
	viperConfigEnvPrefix   = "DOUTOK_USER"
	viperConfigEnvFilename = "user"
)

var (
	Logger = dlog.InitLog(3)
)

func main() {
	config := dtviper.ConfigInit(viperConfigEnvPrefix, viperConfigEnvFilename)

	options, shutdown := initHelper.InitRPCServerArgs(config)
	defer shutdown()

	svr := userservice.NewServer(
		handler.New(),
		options...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
