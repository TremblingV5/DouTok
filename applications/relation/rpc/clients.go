package rpc

import (
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain/messagedomainservice"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain/relationdomainservice"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain/userdomainservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/cloudwego/kitex/client"
)

type Clients struct {
	User     *services.Service[userdomainservice.Client]
	Relation *services.Service[relationdomainservice.Client]
	Message  *services.Service[messagedomainservice.Client]
}

func New(options []client.Option) *Clients {
	return &Clients{
		User:     services.New[userdomainservice.Client](constants.USER_DOMAIN_SERVER_NAME, userdomainservice.NewClient, options),
		Relation: services.New[relationdomainservice.Client](constants.RELATION_DOMAIN_SERVER_NAME, relationdomainservice.NewClient, options),
		Message:  services.New[messagedomainservice.Client](constants.MESSAGE_DOMAIN_SERVER_NAME, messagedomainservice.NewClient, options),
	}
}
