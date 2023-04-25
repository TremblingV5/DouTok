package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

var UserDomainRPCClient *initHelper.UserDomainClient
var RelationDomainRPCClient *initHelper.RelationDomainClient
var MessageDomainRPCClient *initHelper.MessageDomainClient

func Init() {
	UserDomainRPCClient = initHelper.InitUserDomainRPCClient()
	RelationDomainRPCClient = initHelper.InitRelationDomainRPCClient()
	MessageDomainRPCClient = initHelper.InitMessageDomainRPCClient()
}
