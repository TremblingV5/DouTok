package handler

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"strconv"
	"time"
)

func (s *RelationServiceImpl) GetFollowCount(ctx context.Context, req *relation.DouyinRelationCountRequest) (resp *relation.DouyinRelationCountResponse, err error) {
	// TODO: Your code here...
	resp = &relation.DouyinRelationCountResponse{}
	//查缓存
	followKey := utils.KeyGen(req.UserId, 1, 1)
	followerKey := utils.KeyGen(req.UserId, 2, 1)
	followCount, err := redis.RD.Get(ctx, followKey)
	followerCount, err := redis.RD.Get(ctx, followerKey)
	//命中缓存
	if followCount != "" {

		followN, _ := strconv.ParseInt(followCount, 10, 64)
		followerN, _ := strconv.ParseInt(followerCount, 10, 64)
		resp.FollowCount = followN
		resp.FollowerCount = followerN
		resp.StatusCode = 0
		resp.StatusMsg = "success"
		return
	}
	//未命中缓存，查数据库
	count, err := db.QUeryFollowFollowerCOunt(req.UserId)
	if err != nil {
		return
	}
	fmt.Println(count)
	resp.FollowCount = count.FollowCount
	resp.FollowerCount = count.FollowerCount

	//跟新缓存
	followCount = strconv.FormatInt(count.FollowCount, 10)
	followerCount = strconv.FormatInt(count.FollowerCount, 10)
	redis.RD.Set(ctx, followKey, followCount, time.Duration(time.Hour*10))
	redis.RD.Set(ctx, followerKey, followerCount, time.Duration(time.Hour*10))

	return
}
