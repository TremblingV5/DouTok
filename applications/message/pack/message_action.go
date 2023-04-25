package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
)

func PackageMessageActionResponse(result *messageDomain.DoutokAddMessageResponse, err error) (*message.DouyinMessageActionResponse, error) {
	return &message.DouyinMessageActionResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
	}, err
}
