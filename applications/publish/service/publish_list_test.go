package service

import (
	"log"
	"testing"
)

func TestQueryPublishListInHBase(t *testing.T) {
	Init()

	res, err := QueryPublishListInHBase(int64(1111111111111111111))
	if err != nil {
		panic(err)
	}

	log.Println(len(res))

	cnt := 0
	for _, v := range res {
		log.Println(v.GetId())
		cnt++

		if cnt > 10 {
			break
		}
	}
}
