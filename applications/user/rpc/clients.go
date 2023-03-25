package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

var UserDomainRPCClient *initHelper.UserDomainClient
var RelationDomainRPCClient *initHelper.RelationDomainClient

func Init() {
	UserDomainRPCClient = initHelper.InitUserDomainRPCClient()
	RelationDomainRPCClient = initHelper.InitRelationDomainRPCClient()
}
