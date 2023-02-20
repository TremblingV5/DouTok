package service

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestMessageFriendService(t *testing.T) {
	Init()
	msgService := NewMessageFriendService(context.Background())
	req := &message.DouyinFriendListMessageRequest{
		UserId:       10001000,
		FriendIdList: []int64{10002000},
	}
	err, ret := msgService.MessageFriendList(req)
	assert.Nil(t, err)
	for k, v := range ret {
		assert.DeepEqual(t, int64(10002000), k)
		assert.DeepEqual(t, "test msg", v.Content)
	}
}
