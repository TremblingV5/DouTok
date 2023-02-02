package handler

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"testing"
)

func TestRelation(t *testing.T) {
	r := RelationServiceImpl{}

	resp1, err := r.RelationFollowList(context.Background(), &relation.DouyinRelationFollowListRequest{UserId: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp1)

	resp2, err := r.RelationFollowerList(context.Background(), &relation.DouyinRelationFollowerListRequest{UserId: 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp2)
	resp3, err := r.RelationAction(context.Background(), &relation.DouyinRelationActionRequest{UserId: 1, ToUserId: 4, ActionType: 2})
	fmt.Println(resp3)
}
