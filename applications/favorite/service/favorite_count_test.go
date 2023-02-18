package service

import (
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
)

func TestQueryFavoriteCount(t *testing.T) {
	Init()

	video_id := int64(2222222222222222222)
	list := []int64{video_id}

	res, err := QueryFavoriteCount(list)
	if err != nil && err != redis.Nil {
		log.Panicln(err)
	}

	log.Println(res)
}

func TestReadCountFromCache(t *testing.T) {
	Init()
	video_id := int64(1111111111111111111)
	num, ok, err := ReadCountFromCache(video_id)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(num, ok)
}

func TestReadFavTotalCount(t *testing.T) {
	Init()
	video_id := int64(1111111111111111111)
	num, ok, err := ReadFavTotalCount(video_id)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(num, ok)
}
