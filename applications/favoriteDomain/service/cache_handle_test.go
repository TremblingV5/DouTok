package service

import (
	"log"
	"testing"
)

func TestUpdateCacheFavCount(t *testing.T) {
	Init()
	
	video_id := int64(2222222222222222222)

	err1 := UpdateCacheFavCount(video_id, true)
	err2 := UpdateCacheFavCount(video_id, false)

	if err1 != nil || err2 != nil {
		log.Panicln(err1, err2)
	}
}
