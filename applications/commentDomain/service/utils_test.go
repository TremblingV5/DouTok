package service

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetCommentRecall(t *testing.T) {
	timeNow := time.Now().Unix()

	time1 := fmt.Sprint(timeNow - 60*12)
	require.Equal(t, "12分钟前", GetTimeRecall(time1, timeNow))

	time2 := fmt.Sprint(timeNow - 60*60*5)
	require.Equal(t, "5小时前", GetTimeRecall(time2, timeNow))

	time3 := fmt.Sprint(timeNow - 60*60*24*4)
	require.Equal(t, "4天前", GetTimeRecall(time3, timeNow))

	time4 := fmt.Sprint(timeNow - 60*60*24*18)
	require.Equal(t, time.Unix(timeNow-60*60*24*18, 0).Format("2006/01/02"), GetTimeRecall(time4, timeNow))
}
