package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

func TestGetFeedCache(t *testing.T) {
	Init()

	userId := int64(2222222222222222222)
	num := int64(1)

	res, ok := GetFeedCache(context.Background(), fmt.Sprint(userId), num)

	log.Println(ok)
	log.Println(res)
}

func TestSetFeedCache(t *testing.T) {
	Init()

	userId := int64(2222222222222222222)
	method := "l"

	video := VideoInHB{
		Id:         []byte("3333333333333333333"),
		AuthorId:   []byte("2222222222222222222"),
		AuthorName: []byte("Unit testing author name"),
		Title:      []byte("Unit testing title"),
		VideoUrl:   []byte("Unit testing video url"),
		CoverUrl:   []byte("Unit testing cover url"),
		Timestamp:  []byte(fmt.Sprint(time.Now().Unix())),
	}

	err := SetFeedCache(context.Background(), method, fmt.Sprint(userId), video)
	if err != nil {
		log.Panicln(err)
	}
}

func TestGetMarkedTime(t *testing.T) {
	Init()

	userId := int64(2222222222222222222)
	res, err := GetMarkedTime(context.Background(), fmt.Sprint(userId))
	if err != nil && err != redis.Nil {
		panic(err)
	}
	log.Println(res)
}

func TestSetMarkedTime(t *testing.T) {
	Init()

	userId := int64(2222222222222222222)
	err := SetMarkedTime(context.Background(), fmt.Sprint(userId), fmt.Sprint(time.Now().Unix()))

	if err != nil {
		log.Panicln(err)
	}
}
