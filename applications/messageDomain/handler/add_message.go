package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/messageDomain/pack"
	"github.com/TremblingV5/DouTok/applications/messageDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (s *MessageDomainServiceImpl) AddMessage(ctx context.Context, req *messageDomain.DoutokAddMessageRequest) (resp *messageDomain.DoutokAddMessageResponse, err error) {
	resp = new(messageDomain.DoutokAddMessageResponse)

	err = service.NewMessageActionService(ctx).MessageAction(req)
	if err != nil {
		pack.BuildMessageActionResp(err, resp)
		return resp, nil
	}
	pack.BuildMessageActionResp(errno.Success, resp)
	return resp, nil
}
