package service

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestMessageActionService(t *testing.T) {
	Init()
	msgService := NewMessageActionService(context.Background())
	req := &message.DouyinMessageActionRequest{
		UserId:     10001000,
		ToUserId:   10002000,
		ActionType: 1,
		Content:    "test msg",
	}
	err := msgService.MessageAction(req)
	assert.Nil(t, err)
}
