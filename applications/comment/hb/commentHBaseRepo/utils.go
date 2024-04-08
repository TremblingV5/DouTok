package commentHBaseRepo

import (
	"fmt"
	globalMisc "github.com/TremblingV5/DouTok/pkg/misc"
)

func ensureIdLength(id int64) string {
	return globalMisc.LFill(fmt.Sprint(id), 20)
}

func getCommentRowKey(videoId, conversationId int64, isDeleted, timestamp string) string {
	return ensureIdLength(videoId) + isDeleted + ensureIdLength(conversationId) + timestamp
}

func getCommentQueryPrefix(video_id int64) string {
	return ensureIdLength(video_id) + "0"
}
