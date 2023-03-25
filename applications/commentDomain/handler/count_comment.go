package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/applications/commentDomain/pack"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

func (s *CommentDomainServiceImpl) CountComment(ctx context.Context, req *commentDomain.DoutokCountCommentReq) (resp *commentDomain.DoutokCountCommentResp, err error) {
	if len(req.VideoIdList) <= 0 {
		return pack.PackageCountCommentResp(&misc.SystemErr, nil)
	}

	res, errNo, err := service.NewCountCommentService(ctx).CountComment(req.VideoIdList...)
	if err != nil {
		return pack.PackageCountCommentResp(errNo, nil)
	}

	return pack.PackageCountCommentResp(&misc.Success, res)
}
