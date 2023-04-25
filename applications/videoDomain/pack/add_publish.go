package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageAddPublishResp(errNo *errno.ErrNo) (request *videoDomain.DoutokAddPublishResponse, err error) {
	return &videoDomain.DoutokAddPublishResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
	}, nil
}
