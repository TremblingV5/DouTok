package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (s *RelationDomainServiceImpl) AddRelation(ctx context.Context, req *relationDomain.DoutokAddRelationRequest) (resp *relationDomain.DoutokAddRelationResponse, err error) {
	resp = new(relationDomain.DoutokAddRelationResponse)

	err = service.NewRelationActionService(ctx).AddRelation(req)
	if err != nil {
		pack.BuildRelationActionResp(err, resp)
		return resp, nil
	}
	pack.BuildRelationActionResp(errno.Success, resp)
	return resp, nil
}
