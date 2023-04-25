package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (s *RelationDomainServiceImpl) RmRelation(ctx context.Context, req *relationDomain.DoutokRmRelationRequest) (resp *relationDomain.DoutokRmRelationResponse, err error) {
	resp = new(relationDomain.DoutokRmRelationResponse)

	err = service.NewRelationActionService(ctx).RmRelation(req)
	if err != nil {
		pack.BuildRmRelationActionResp(err, resp)
		return resp, nil
	}
	pack.BuildRmRelationActionResp(errno.Success, resp)
	return resp, nil
}
