package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

func (s *CommentDomainHandler) CountComment(ctx context.Context, req *commentDomain.DoutokCountCommentReq) (resp *commentDomain.DoutokCountCommentResp, err error) {
	if len(req.VideoIdList) <= 0 {
		return &commentDomain.DoutokCountCommentResp{
			StatusCode:   1,
			StatusMsg:    ParametersError.Error(),
			CommentCount: nil,
		}, ParametersError
	}

	result, err := s.service.CountComments(ctx, req.VideoIdList...)
	if err != nil {
		return &commentDomain.DoutokCountCommentResp{
			StatusCode:   1,
			StatusMsg:    err.Error(),
			CommentCount: nil,
		}, err
	}

	return &commentDomain.DoutokCountCommentResp{
		StatusCode:   0,
		StatusMsg:    Success,
		CommentCount: result,
	}, nil
}
