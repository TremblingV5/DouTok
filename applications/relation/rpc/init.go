package rpc

import "github.com/TremblingV5/DouTok/pkg/dtviper"

// InitRPC init rpc client
func InitRPC() {
	MessageConfig := dtviper.ConfigInit("DOUTOK_MESSAGE", "message")
	initMessageRpc(&MessageConfig)

	UserConfig := dtviper.ConfigInit("DOUTOK_USER", "user")
	initUserRpc(&UserConfig)
}
