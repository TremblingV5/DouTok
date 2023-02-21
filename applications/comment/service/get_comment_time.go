package service

import (
	"fmt"
	"strconv"
	"time"
)

func GetTimeRecall(timestamp string, current int64) string {
	timestampI64, _ := strconv.ParseInt(timestamp, 10, 64)

	diff := current - timestampI64

	if diff < 60 {
		return "刚刚"
	} else if diff < 60*60 {
		minutes := diff / 60
		return fmt.Sprint(minutes) + "分钟前"
	} else if diff < 60*60*24 {
		hours := diff / (60 * 60)
		return fmt.Sprint(hours) + "小时前"
	} else if diff < 60*60*24*14 {
		days := diff / (60 * 60 * 24)
		return fmt.Sprint(days) + "天前"
	} else {
		t := time.Unix(timestampI64, 0)
		return t.Format("2006/01/02")
	}
}
