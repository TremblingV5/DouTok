package service

import (
	"log"
	"testing"
)

func TestWriteFavoriteInCache(t *testing.T) {
	Init()

	userId := int64(1111111111111111111)
	videoId := int64(2222222222222222222)

	err := WriteFavoriteInCache(userId, videoId, true)
	if err != nil {
		log.Panicln(err)
	}
}
