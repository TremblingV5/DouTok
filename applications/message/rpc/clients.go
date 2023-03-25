package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

var MessageDomainClient *initHelper.MessageDomainClient

func Init() {
	MessageDomainClient = initHelper.InitMessageDomainRPCClient()
}
