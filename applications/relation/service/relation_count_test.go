package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestRelationCountService(t *testing.T) {
	Init()
	relService := NewRelationCountService(context.Background())
	req := &relation.DouyinRelationCountRequest{
		UserId: 10001000,
	}
	err, follow, follower := relService.RelationCount(req)
	assert.Nil(t, err)
	fmt.Printf("userId = 10001000, follow = %d, follower = %d\n", follow, follower)

	relService = NewRelationCountService(context.Background())
	req = &relation.DouyinRelationCountRequest{
		UserId: 10004000,
	}
	err, follow, follower = relService.RelationCount(req)
	assert.Nil(t, err)
	fmt.Printf("userId = 10004000, follow = %d, follower = %d\n", follow, follower)
}
