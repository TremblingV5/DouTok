package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

var FavoriteDomainClient *initHelper.FavoriteDomainClient
var VideoDomainClient *initHelper.VideoDomainClient

func Init() {
	FavoriteDomainClient = initHelper.InitFavoriteDomainRPCClient()
	VideoDomainClient = initHelper.InitVideoDomainRPCClient()
}
