package rpc

import "github.com/TremblingV5/DouTok/pkg/initHelper"

type Clients struct {
	User     *initHelper.UserDomainClient
	Relation *initHelper.RelationDomainClient
}

func New() *Clients {
	return &Clients{
		User:     initHelper.InitUserDomainRPCClient(),
		Relation: initHelper.InitRelationDomainRPCClient(),
	}
}
