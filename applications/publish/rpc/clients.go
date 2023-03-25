package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

var UserDomainClient *initHelper.UserDomainClient
var VideoDomainClient *initHelper.VideoDomainClient
var FavoriteDomainClient *initHelper.FavoriteDomainClient

func Init() {
	UserDomainClient = initHelper.InitUserDomainRPCClient()
	VideoDomainClient = initHelper.InitVideoDomainRPCClient()
	FavoriteDomainClient = initHelper.InitFavoriteDomainRPCClient()
}
