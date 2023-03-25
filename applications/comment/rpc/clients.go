package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

var UserDomainRPCClient *initHelper.UserDomainClient
var CommentDomainRPCClient *initHelper.CommentDomainClient

func Init() {
	UserDomainRPCClient = initHelper.InitUserDomainRPCClient()
	CommentDomainRPCClient = initHelper.InitCommentDomainRPCClient()
}
