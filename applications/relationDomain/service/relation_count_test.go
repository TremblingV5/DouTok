package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestRelationCountService(t *testing.T) {
	Init()
	relService := NewRelationCountService(context.Background())
	// req := &relationDomain.DoutokCountRelationRequest{
	// 	UserId: []int64{10001000},
	// }
	err, follow, follower := relService.RelationCount(10001000)
	assert.Nil(t, err)
	fmt.Printf("userId = 10001000, follow = %d, follower = %d\n", follow, follower)

	relService = NewRelationCountService(context.Background())
	// req = &relation.DouyinRelationCountRequest{
	// 	UserId: 10004000,
	// }
	err, follow, follower = relService.RelationCount(10004000)
	assert.Nil(t, err)
	fmt.Printf("userId = 10004000, follow = %d, follower = %d\n", follow, follower)
}
