package pack

import "github.com/TremblingV5/DouTok/kitex_gen/publish"

func PackPublishActionRes(code int32, msg string) (*publish.DouyinPublishActionResponse, error) {
	var resp publish.DouyinPublishActionResponse

	resp.StatusCode = code
	resp.StatusMsg = msg

	return &resp, nil
}
