package service

import (
	"fmt"
)

/*
	op为true表示增加1个喜欢数，否则表示减少1个喜欢数
*/
func UpdateCacheFavCount(video_id int64, op bool) error {
	data, ok := FavCount.Get(fmt.Sprint(video_id))

	if ok {
		if op {
			FavCount.Set(fmt.Sprint(video_id), data.(int)+1)
		} else {
			FavCount.Set(fmt.Sprint(video_id), data.(int)-1)
		}
	} else {
		if op {
			FavCount.Set(fmt.Sprint(video_id), 1)
		} else {
			FavCount.Set(fmt.Sprint(video_id), -1)
		}
	}

	return nil
}
