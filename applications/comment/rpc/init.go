package rpc

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func InitRPCConfig() error {
	var config configStruct.CommentConfig
	configurator.InitConfig(
		&config, "comment.yaml",
	)

	ClientConfig = &config

	return nil
}

func InitPRCClient() error {
	InitUserRpc()

	return nil
}
