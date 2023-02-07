package handler

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"testing"
)

func TestRelation(t *testing.T) {

	//读取配置
	v, err := conf.InitConfig("./", "relation")
	if err != nil {
		return
	}
	//连接数据库
	db.Conn(v)
	//连接redis
	redis.Conn(v)

	r := RelationServiceImpl{}
	resp3, err := r.RelationAction(context.Background(), &relation.DouyinRelationActionRequest{UserId: 1, ToUserId: 2, ActionType: 1})
	fmt.Println("关注结果", resp3)

	resp1, err := r.RelationFollowList(context.Background(), &relation.DouyinRelationFollowListRequest{UserId: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println("关注列表：", resp1)

	resp2, err := r.RelationFollowerList(context.Background(), &relation.DouyinRelationFollowerListRequest{UserId: 2})
	if err != nil {
		panic(err)
	}
	fmt.Println("粉丝列表：", resp2)
}
