package handler

import (
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
)

type RelationDomainServiceImpl struct {
}

type RelationDomainHandler struct {
	service *service.RelationDomainService
}

func NewRelationDomainHandler(service *service.RelationDomainService) *RelationDomainHandler {
	return &RelationDomainHandler{
		service: service,
	}
}
