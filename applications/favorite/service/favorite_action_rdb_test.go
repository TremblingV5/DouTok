package service

import (
	"log"
	"testing"
)

func TestCreateFavoriteInRDB(t *testing.T) {
	Init()

	userId := int64(1111111111111111111)
	videoId := int64(2222222222222222222)
	op := true

	err := CreateFavoriteInRDB(userId, videoId, op)
	if err != nil {
		log.Panicln(err)
	}
}
