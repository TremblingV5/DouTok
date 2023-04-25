package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageListFavoriteResp(errNo *errno.ErrNo, data []int64) (resp *favoriteDomain.DoutokListFavResponse, err error) {
	var result []*entity.Video

	for _, v := range data {
		result = append(result, &entity.Video{
			Id: v,
		})
	}

	return &favoriteDomain.DoutokListFavResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
		VideoList:  result,
	}, nil
}
