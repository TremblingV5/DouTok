package rpc

import (
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/cloudwego/kitex/client"

	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain/relationdomainservice"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain/userdomainservice"
	"github.com/TremblingV5/DouTok/pkg/services"
)

type Clients struct {
	User     *services.Service[userdomainservice.Client]
	Relation *services.Service[relationdomainservice.Client]
}

func New(options []client.Option) *Clients {
	return &Clients{
		User:     services.New[userdomainservice.Client](constants.USER_DOMAIN_SERVER_NAME, userdomainservice.NewClient, options),
		Relation: services.New[relationdomainservice.Client](constants.RELATION_DOMAIN_SERVER_NAME, relationdomainservice.NewClient, options),
	}
}
