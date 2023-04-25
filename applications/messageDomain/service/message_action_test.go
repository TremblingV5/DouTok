package service

import (
	"context"
	"testing"

	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestMessageActionService(t *testing.T) {
	Init()
	msgService := NewMessageActionService(context.Background())
	req := &messageDomain.DoutokAddMessageRequest{
		UserId:     10001000,
		ToUserId:   10002000,
		ActionType: 1,
		Content:    "test msg",
	}
	err := msgService.MessageAction(req)
	assert.Nil(t, err)
}
