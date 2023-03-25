package pack

import (
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/model"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageListCommentResp(errNo *errno.ErrNo, data []*model.CommentInHB) (resp *commentDomain.DoutokListCommentResp, err error) {
	result := []*entity.Comment{}

	for _, v := range data {
		result = append(result, &entity.Comment{
			Id: v.GetId(),
			User: &entity.User{
				Id: v.GetUserId(),
			},
			Content:    v.GetContent(),
			CreateDate: v.GetTimestamp(),
		})
	}

	return &commentDomain.DoutokListCommentResp{
		StatusCode:  int32(errNo.ErrCode),
		StatusMsg:   errNo.ErrMsg,
		CommentList: result,
	}, nil
}
