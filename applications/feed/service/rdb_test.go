package service

import (
	"log"
	"testing"
)

func TestGetVideoByIdInRDB(t *testing.T) {
	Init()

	res, err := GetVideoByIdInRDB(uint64(1627183461678981120))
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res)
}
