package redishandle

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"testing"
)

func TestSGet(t *testing.T) {
	redisConfig := configStruct.RedisConfig{}
	err := configurator.InitConfig(&redisConfig, "relation_redis.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	clientCache, err := InitRedis(redisConfig.Host+":"+redisConfig.Port, redisConfig.Password, redisConfig.Databases)
	if err != nil {
		fmt.Println(err)
		return
	}

	rd := clientCache[constants.DbDefault]
	res, err := rd.Client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	k := utils.KeyGen(1, 1, 2)
	fmt.Println(k)
	if r, err := rd.SGet(context.Background(), k); err != nil {
		panic(err)
	} else {
		fmt.Println(r)
	}
	if r, err := rd.SGetObj(context.Background(), k); err != nil {
		panic(err)
	} else {
		fmt.Println(r)
	}
}
