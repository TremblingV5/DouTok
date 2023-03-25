package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageCountFavoriteResp(errNo *errno.ErrNo, data map[int64]int64) (resp *favoriteDomain.DoutokCountFavResponse, err error) {
	return &favoriteDomain.DoutokCountFavResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
		CountFav:   data,
	}, nil
}
