package service

import (
	"fmt"
	"testing"
	"time"
)

func TestSavePublish(t *testing.T) {
	Init()

	timestamp := fmt.Sprint(time.Now().Unix())

	userId := int64(1111111111111111111)
	title := "Unit testing on " + timestamp
	file := []byte("Unit testing on " + timestamp)

	err := SavePublish(userId, title, file)

	if err != nil {
		panic(err)
	}
}
