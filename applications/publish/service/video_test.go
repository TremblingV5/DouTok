package service

import (
	"fmt"
	"testing"
)

func TestQueryVideoCountFromRDB(t *testing.T) {
	Init()

	v, err := QueryVideoFromRBDById(1627179313688485888)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(v)
}

func TestQueryAFewVideoCountFromRDB(t *testing.T) {
	Init()

	videos, err := QuerySomeVideoFromRDBByIds(1627179313688485888, 1627180724060954624)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(videos)
}
