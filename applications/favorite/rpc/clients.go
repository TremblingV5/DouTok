package rpc

import (
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain/favoritedomainservice"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain/videodomainservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"github.com/cloudwego/kitex/client"
)

type Clients struct {
	Favorite *services.Service[favoritedomainservice.Client]
	Video    *services.Service[videodomainservice.Client]
}

func New(options []client.Option) *Clients {
	return &Clients{
		Favorite: services.New[favoritedomainservice.Client](constants.FAVORITE_DOMAIN_SERVER_NAME, favoritedomainservice.NewClient, options),
		Video:    services.New[videodomainservice.Client](constants.VIDEO_DOMAIN_SERVER_NAME, videodomainservice.NewClient, options),
	}
}
