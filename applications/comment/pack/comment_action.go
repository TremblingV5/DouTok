package pack

import "github.com/TremblingV5/DouTok/kitex_gen/comment"

func PackCommentActionResp(code int32, msg string, cmt *comment.Comment) (resp *comment.DouyinCommentActionResponse, err error) {
	return &comment.DouyinCommentActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Comment:    cmt,
	}, nil
}
