package service

import (
	"log"
	"testing"
)

func TestListComment(t *testing.T) {
	Init()

	videoId := int64(1111111111111111111)
	userId := int64(1111111111111111111)
	conId := int64(1111111111111111111)
	lastId := int64(0)
	toUserId := int64(2222222222222222)
	content := "Unit test comment content"

	res, err := AddComment(
		videoId, userId, conId, lastId, toUserId, content,
	)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res)

	videoList, errNo, err := ListComment(videoId)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(errNo)
	log.Println(videoList)
}
