package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func PackagePublishActionResponse(result *videoDomain.DoutokAddPublishResponse, err error) (*publish.DouyinPublishActionResponse, error) {
	if err != nil {
		return nil, err
	}

	return &publish.DouyinPublishActionResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
	}, err
}
