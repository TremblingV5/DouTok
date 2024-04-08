package service

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestFindFeedInHB(t *testing.T) {
	Init()

	res, err := FindFeedInHB("1627500000", fmt.Sprint(time.Now().Unix()))
	if err != nil {
		log.Panicln(err)
	}

	log.Println(len(res))
}

func TestSearchFeedEarlierInHB(t *testing.T) {
	Init()

	curr := time.Now().Unix()

	res, err := SearchFeedEarlierInHB(curr, curr-86400*7)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(len(res))
}

func TestSearchFeedLaterInHB(t *testing.T) {
	Init()

	curr := time.Now().Unix()

	res, newMarkedTime, err := SearchFeedLaterInHB(fmt.Sprint(curr-86400), fmt.Sprint(curr))
	if err != nil {
		log.Panicln(err)
	}

	log.Println(len(res))
	log.Println(newMarkedTime)
}
