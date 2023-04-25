package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageAddCommentResp(errNo *errno.ErrNo, data *entity.Comment, userId int64) (resp *commentDomain.DoutokAddCommentResp, err error) {
	return &commentDomain.DoutokAddCommentResp{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
		Comment:    data,
		//Comment: &entity.Comment{
		//	Id: data.Id,
		//	User: &entity.User{
		//		Id: data.UserId,
		//	},
		//	Content:    data.Content,
		//	CreateDate: data.Timestamp,
		//},
	}, nil
}
