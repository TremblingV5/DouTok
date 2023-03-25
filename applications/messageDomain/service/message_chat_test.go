package service

import (
	"context"
	"testing"

	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestMessageChatService(t *testing.T) {
	Init()
	msgService := NewMessageChatService(context.Background())
	req := &messageDomain.DoutokListMessageRequest{
		ToUserId:   10002000,
		UserId:     10001000,
		PreMsgTime: 0,
	}
	err, ret := msgService.MessageChat(req)
	assert.Nil(t, err)
	for _, msg := range ret {
		assert.DeepEqual(t, "test msg", msg.Content)
	}
}
