package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/misc"
	"github.com/TremblingV5/DouTok/applications/commentDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
)

func (s *CommentDomainHandler) ListComment(ctx context.Context, req *commentDomain.DoutokListCommentReq) (resp *commentDomain.DoutokListCommentResp, err error) {
	if !misc.CheckCommentListArgs(req) {
		return &commentDomain.DoutokListCommentResp{
			StatusCode:  1,
			StatusMsg:   ParametersError.Error(),
			CommentList: nil,
		}, ParametersError
	}

	list, err := service.DomainUtil.ListComment(ctx, req.VideoId)
	if err != nil {
		return &commentDomain.DoutokListCommentResp{
			StatusCode:  1,
			StatusMsg:   err.Error(),
			CommentList: nil,
		}, err
	}

	result := func() []*entity.Comment {
		var r []*entity.Comment
		for _, v := range list {
			r = append(r, &entity.Comment{
				Id: v.GetId(),
				User: &entity.User{
					Id: v.GetUserId(),
				},
				Content:    v.GetContent(),
				CreateDate: v.GetTimestamp(),
			})
		}
		return r
	}()

	return &commentDomain.DoutokListCommentResp{
		StatusCode:  0,
		StatusMsg:   Success,
		CommentList: result,
	}, nil
}
