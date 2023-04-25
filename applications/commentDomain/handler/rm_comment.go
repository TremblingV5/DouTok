package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/applications/commentDomain/pack"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

func (s *CommentDomainServiceImpl) RmComment(ctx context.Context, req *commentDomain.DoutokRmCommentReq) (resp *commentDomain.DoutokAddCommentResp, err error) {
	errNo, err := service.NewRmCommentService(ctx).RmComment(req.UserId, req.CommentId)
	if err != nil {
		return pack.PackageAddCommentResp(&errNo, nil, req.UserId)
	}

	return pack.PackageAddCommentResp(&misc.Success, nil, req.UserId)
}
