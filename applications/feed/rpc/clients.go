package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

var CommentDomainClient *initHelper.CommentDomainClient
var FavoriteDomainClient *initHelper.FavoriteDomainClient
var UserDomainClient *initHelper.UserDomainClient
var VideoDomainClient *initHelper.VideoDomainClient

func Init() {
	CommentDomainClient = initHelper.InitCommentDomainRPCClient()
	FavoriteDomainClient = initHelper.InitFavoriteDomainRPCClient()
	UserDomainClient = initHelper.InitUserDomainRPCClient()
	VideoDomainClient = initHelper.InitVideoDomainRPCClient()
}
