package rpc

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func InitRPCConfig() error {
	var config configStruct.FavoriteConfig
	configurator.InitConfig(
		&config, "user.yaml",
	)

	ClientConfig = &config

	return nil
}

func InitPRCClient() error {
	InitRelationRpc()

	return nil
}
