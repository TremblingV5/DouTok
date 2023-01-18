package service

import (
	"fmt"
	"testing"
)

func TestQueryVideoCountFromRDB(t *testing.T) {
	GetDb()

	v, err := QueryVideoCountFromRDB(1)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(v)
}

func TestQueryAFewVideoCountFromRDB(t *testing.T) {
	GetDb()

	videos, err := QueryAFewVideoCountFromRDB(1, 22)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(videos)
}
