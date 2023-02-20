package service

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestJudgeTimeDiff(t *testing.T) {
	curr := time.Now().Unix()
	ok := JudgeTimeDiff(curr-86400, fmt.Sprint(curr), 8000)

	if ok {
		log.Panicln(ok)
	}
}

func TestJudgeQuantityRatio(t *testing.T) {
	ok := JudgeQuantityRatio(2, 3, 0.5)

	if !ok {
		log.Panicln(ok)
	}
}
