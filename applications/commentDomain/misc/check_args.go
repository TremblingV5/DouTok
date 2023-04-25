package misc

import (
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
)

func CheckCommentActionArgs(req *commentDomain.DoutokAddCommentReq) bool {
	return req.CommentText != ""
}

func CheckCommentListArgs(req *commentDomain.DoutokListCommentReq) bool {
	return req.VideoId >= 0
}
