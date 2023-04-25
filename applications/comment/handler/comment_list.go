package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"

	"github.com/TremblingV5/DouTok/applications/comment/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
)

func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	result, err := rpc.CommentDomainRPCClient.ListComment(ctx, &commentDomain.DoutokListCommentReq{
		VideoId: req.VideoId,
	})

	return pack.PackageCommentListRepsonse(ctx, result, err)
}
