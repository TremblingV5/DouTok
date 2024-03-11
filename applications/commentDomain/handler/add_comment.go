package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

func (s *CommentDomainHandler) AddComment(ctx context.Context, req *commentDomain.DoutokAddCommentReq) (resp *commentDomain.DoutokAddCommentResp, err error) {
	if ok := misc.CheckCommentActionArgs(req); !ok {
		return &commentDomain.DoutokAddCommentResp{
			StatusCode: 1,
			StatusMsg:  ParametersError.Error(),
			Comment:    nil,
		}, ParametersError
	}

	result, err := service.DomainUtil.AddComment(ctx, req.VideoId, req.UserId, service.DomainUtil.SnowflakeHandle.GetId().Int64(), 0, 0, req.CommentText)
	if err != nil {
		return &commentDomain.DoutokAddCommentResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			Comment:    nil,
		}, err
	}

	return &commentDomain.DoutokAddCommentResp{
		StatusCode: 0,
		StatusMsg:  Success,
		Comment:    result,
	}, nil
}
