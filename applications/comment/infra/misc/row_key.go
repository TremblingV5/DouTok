package misc

import (
	"fmt"

	globalMisc "github.com/TremblingV5/DouTok/pkg/misc"
)

func EnsureIdLength(id int64) string {
	return globalMisc.LFill(fmt.Sprint(id), 20)
}

func GetCommentRowKey(video_id int64, is_deleted string, conversation_id int64, timestamp string) string {
	return EnsureIdLength(video_id) + is_deleted + EnsureIdLength(conversation_id) + timestamp
}

func GetCommentQueryPrefix(video_id int64) string {
	return EnsureIdLength(video_id) + "0"
}
