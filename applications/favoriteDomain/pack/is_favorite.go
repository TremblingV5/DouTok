package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageIsFavoriteResp(errNo *errno.ErrNo, data map[int64]bool) (resp *favoriteDomain.DoutokIsFavResponse, err error) {
	return &favoriteDomain.DoutokIsFavResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
		IsFav:      data,
	}, nil
}
