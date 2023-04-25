package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
)

func TestHBaseType2RPCType(t *testing.T) {
	video := typedef.VideoInHB{
		Id:       []byte(fmt.Sprint(uint64(123456789123456789))),
		AuthorId: []byte(fmt.Sprint(uint64(987654321987654321))),
		Title:    []byte("title"),
	}
	log.Println(video)

	res := HBaseType2RPCType(video)
	log.Println(res)
}
