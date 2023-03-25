package service

import (
	"context"
	"log"
	"testing"
)

func TestActionFavorite(t *testing.T) {
	Init()

	userId := int64(1111111111111111111)
	videoId := int64(2222222222222222222)

	errNo, err := NewActionFavoriteService(context.Background()).ActionFavorite(userId, videoId, true)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(errNo)

	errNo, err = NewActionFavoriteService(context.Background()).ActionFavorite(userId, videoId, false)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(errNo)
}
