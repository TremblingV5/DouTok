package rpc

import (
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain/commentdomainservice"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain/userdomainservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/cloudwego/kitex/client"
)

type Clients struct {
	User    *services.Service[userdomainservice.Client]
	Comment *services.Service[commentdomainservice.Client]
}

func New(options []client.Option) *Clients {
	return &Clients{
		User:    services.New[userdomainservice.Client](constants.USER_DOMAIN_SERVER_NAME, userdomainservice.NewClient, options),
		Comment: services.New[commentdomainservice.Client](constants.COMMENT_DOMAIN_SERVER_NAME, commentdomainservice.NewClient, options),
	}
}
