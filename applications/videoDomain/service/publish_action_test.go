package service

import (
	"context"
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

	err := NewSavePublishService(context.Background()).SavePublish(userId, title, file)

	if err != nil {
		panic(err)
	}
}
