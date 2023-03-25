package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/applications/commentDomain/pack"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

func (s *CommentDomainServiceImpl) AddComment(ctx context.Context, req *commentDomain.DoutokAddCommentReq) (resp *commentDomain.DoutokAddCommentResp, err error) {
	if result := misc.CheckCommentActionArgs(req); !result {
		return pack.PackageAddCommentResp(&misc.ParamsErr, nil, req.UserId)
	}

	result, err := service.NewAddCommentService(ctx).AddComment(req.VideoId, req.UserId, utils.GetSnowFlakeId().Int64(), 0, 0, req.CommentText)
	if err != nil {
		return pack.PackageAddCommentResp(&misc.SystemErr, nil, req.UserId)
	}

	return pack.PackageAddCommentResp(&misc.Success, result, req.UserId)
}
