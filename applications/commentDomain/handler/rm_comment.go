package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

func (s *CommentDomainHandler) RmComment(ctx context.Context, req *commentDomain.DoutokRmCommentReq) (resp *commentDomain.DoutokAddCommentResp, err error) {
	if err := service.DomainUtil.RemoveComment(ctx, req.UserId, req.CommentId); err != nil {
		return &commentDomain.DoutokAddCommentResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			Comment:    nil,
		}, err
	}

	return &commentDomain.DoutokAddCommentResp{
		StatusCode: 0,
		StatusMsg:  Success,
		Comment:    nil,
	}, nil
}
