package redis

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/constants"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/spf13/viper"
)

var RD *redishandle.RedisClient

func Conn(v *viper.Viper) error {
	host := v.GetString("redis.host")
	port := v.GetString("redis.port")
	password := v.GetString("redis.password")
	d := v.GetStringMap("redis.databases")
	database := make(map[string]int)
	//fmt.Println(host, port, d)
	for k, val := range d {
		database[k] = val.(int)
	}
	clientCache, err := redishandle.InitRedis(host+":"+port, password, database)
	if err != nil {
		return err
	}
	RD = clientCache[constants.DbDefault]
	r, _ := RD.Client.Ping(context.Background()).Result()
	fmt.Println(r)
	return nil
}

func Keys(userId, toUserId int64) []string {
	r := make([]string, 4)
	r[0] = utils.KeyGen(userId, 1, 1)
	r[1] = utils.KeyGen(userId, 1, 2)
	r[2] = utils.KeyGen(toUserId, 2, 1)
	r[3] = utils.KeyGen(toUserId, 2, 2)
	return r
}
