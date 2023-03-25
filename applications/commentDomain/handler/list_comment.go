package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/applications/commentDomain/pack"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

func (s *CommentDomainServiceImpl) ListComment(ctx context.Context, req *commentDomain.DoutokListCommentReq) (resp *commentDomain.DoutokListCommentResp, err error) {
	if !misc.CheckCommentListArgs(req) {
		return pack.PackageListCommentResp(&misc.VideoIdErr, nil)
	}

	res, errNo, err := service.NewListCommentService(ctx).ListComment(req.VideoId)
	if err != nil {
		return pack.PackageListCommentResp(&errNo, nil)
	}

	return pack.PackageListCommentResp(&misc.Success, res)
}
