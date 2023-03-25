package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageCountCommentResp(errNo *errno.ErrNo, data map[int64]int64) (resp *commentDomain.DoutokCountCommentResp, err error) {
	return &commentDomain.DoutokCountCommentResp{
		StatusCode:   int32(errNo.ErrCode),
		StatusMsg:    errNo.ErrMsg,
		CommentCount: data,
	}, nil
}
