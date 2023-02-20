package service

import (
	"log"
	"testing"
)

func TestQueryIsFavorite(t *testing.T) {
	Init()

	res, err := QueryIsFavorite(int64(1111111111111111111), []int64{int64(2222222222222222222)})
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res)
}
