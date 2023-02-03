package rpc

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func InitRPCConfig() error {
	var config configStruct.MessageConfig
	configurator.InitConfig(
		&config, "message.yaml",
	)

	ClientConfig = &config

	return nil
}
