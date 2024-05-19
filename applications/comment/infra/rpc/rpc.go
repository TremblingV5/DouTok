package rpc

import (
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/cloudwego/kitex/client"
)

type Clients struct {
	User *services.Service[userservice.Client]
}

func New(options []client.Option) *Clients {
	return &Clients{
		User: services.New[userservice.Client](constants.UserServerName, userservice.NewClient, options),
	}
}
