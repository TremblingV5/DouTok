package service

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGetCommentRecall(t *testing.T) {
	timeNow := time.Now().Unix()

	// 1小时内的时间戳，返回 x分钟前
	time1 := fmt.Sprint(timeNow - 60*12)
	log.Println(GetTimeRecall(time1, timeNow))

	// 1天内的时间戳，返回 x小时前
	time2 := fmt.Sprint(timeNow - 60*60*5)
	log.Println(GetTimeRecall(time2, timeNow))

	// 14天内的时间戳，返回 x天前
	time3 := fmt.Sprint(timeNow - 60*60*24*4)
	log.Println(GetTimeRecall(time3, timeNow))

	// 14天外的时间戳，返回具体日期
	time4 := fmt.Sprint(timeNow - 60*60*24*18)
	log.Println(GetTimeRecall(time4, timeNow))
}
