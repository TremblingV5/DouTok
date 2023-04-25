package service

import (
	"testing"
	"time"
)

func TestConsumeMsg(t *testing.T) {
	Init()
	go ConsumeMsg()
	time.Sleep(time.Second * 3)
}

// func TestFlush(t *testing.T) {
// 	Init()

// 	relService := NewRelationActionService(context.Background())
// 	req := &relationDomain.DoutokAddRelationRequest{
// 		UserId:   10001000,
// 		ToUserId: 10002000,
// 		// ActionType: 1,
// 	}
// 	err := relService.RelationAction(req)
// 	assert.Nil(t, err)

// 	go Flush()
// 	time.Sleep(time.Second * 3)
// }
