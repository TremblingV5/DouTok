package misc

import "github.com/TremblingV5/DouTok/pkg/misc"

func FillUserId(user_id string) string {
	return misc.LFill(user_id, 20)
}
