package service

import (
	"fmt"
	"testing"
)

func TestQueryVideoCountFromRDB(t *testing.T) {
	InitDb()

	v, err := QueryVideoFromRBDById(1)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(v)
}

func TestQueryAFewVideoCountFromRDB(t *testing.T) {
	InitDb()

	videos, err := QuerySomeVideoFromRDBByIds(1, 22)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(videos)
}
