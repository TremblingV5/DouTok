package misc

import "github.com/TremblingV5/DouTok/kitex_gen/comment"

func CheckCommentActionArgs(req *comment.DouyinCommentActionRequest) bool {
	return req.CommentText != ""
}

func CheckCommentListArgs(req *comment.DouyinCommentListRequest) bool {
	return req.VideoId >= 0
}
