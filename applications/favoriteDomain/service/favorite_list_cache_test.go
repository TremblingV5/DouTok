package service

import (
	"log"
	"testing"
)

func TestQueryFavListInCache(t *testing.T) {
	Init()

	userId := int64(1111111111111111111)
	res, err := QueryFavListInCache(userId)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res)
}

func TestWriteFavListInCache(t *testing.T) {
	Init()

	userId := int64(1111111111111111111)
	list := []int64{
		int64(2222222222222222222), int64(2222222222222222223),
	}
	err := WriteFavListInCache(userId, list)

	if err != nil {
		log.Panicln(err)
	}
}
