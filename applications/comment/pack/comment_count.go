package pack

import "github.com/TremblingV5/DouTok/kitex_gen/comment"

func PackCommentCountResp(code int32, msg string, countList map[int64]int64) (*comment.DouyinCommentCountResponse, error) {
	return &comment.DouyinCommentCountResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Result:     countList,
	}, nil
}
