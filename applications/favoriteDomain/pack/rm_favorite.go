package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageRmFavoriteResp(errNo *errno.ErrNo) (resp *favoriteDomain.DoutokRmFavResponse, err error) {
	return &favoriteDomain.DoutokRmFavResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
	}, nil
}
