package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageAddFavoriteResp(errNo *errno.ErrNo) (resp *favoriteDomain.DoutokAddFavResponse, err error) {
	return &favoriteDomain.DoutokAddFavResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
	}, nil
}
