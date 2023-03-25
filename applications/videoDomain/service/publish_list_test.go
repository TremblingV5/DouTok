package service

import (
	"context"
	"log"
	"testing"
)

func TestQueryPublishListInHBase(t *testing.T) {
	Init()

	res, err := NewQueryPublishListService(context.Background()).QueryPublishListInHBase(int64(1650819877570772992))
	if err != nil {
		panic(err)
	}

	log.Println(len(res))

	cnt := 0
	for _, v := range res {
		log.Println(v.GetId())
		log.Println(v.Id)
		cnt++

		if cnt > 10 {
			break
		}
	}
}
