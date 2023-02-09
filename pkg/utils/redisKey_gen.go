package utils

import (
	"github.com/TremblingV5/DouTok/pkg/constants"
	"strconv"
	"strings"
)

func KeyGen(userId int64, followOrfollower, countOrlist int) string {
	u := strconv.FormatInt(userId, 10)
	if followOrfollower == 1 {
		if countOrlist == 1 {
			return strings.Join([]string{u, "follow", "count"}, constants.KeySep)
		}
		return strings.Join([]string{u, "follow", "list"}, constants.KeySep)
	} else {
		if countOrlist == 1 {
			return strings.Join([]string{u, "follower", "count"}, constants.KeySep)
		}
		return strings.Join([]string{u, "follower", "list"}, constants.KeySep)
	}
}
