package rpc

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func InitRPCConfig() error {
	var config configStruct.FeedConfig
	configurator.InitConfig(
		&config, "feed.yaml",
	)

	ClientConfig = &config

	return nil
}

func InitPRCClient() error {
	InitCommentRpc()
	InitFavoriteRpc()
	InitUserRpc()

	return nil
}
