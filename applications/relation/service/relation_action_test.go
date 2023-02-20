package service

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestRelationActionService(t *testing.T) {
	Init()
	relService := NewRelationActionService(context.Background())
	req := &relation.DouyinRelationActionRequest{
		UserId:     10001000,
		ToUserId:   10002000,
		ActionType: 1,
	}
	err := relService.RelationAction(req)
	assert.Nil(t, err)
}
