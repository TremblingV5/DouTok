package redis

import (
	"fmt"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

var RD *redishandle.RedisClient

func Conn() {
	redisConfig := configStruct.RedisConfig{}
	err := configurator.InitConfig(&redisConfig, "relation_redis.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	clientCache, err := redishandle.InitRedis(redisConfig.Host+":"+redisConfig.Port, redisConfig.Password, redisConfig.Databases)
	if err != nil {
		fmt.Println(err)
		return
	}
	RD = clientCache[constants.DbDefault]

}

func init() {
	Conn()
}

func Keys(userId, toUserId int64) []string {
	r := make([]string, 4)
	r[0] = utils.KeyGen(userId, 1, 1)
	r[1] = utils.KeyGen(userId, 1, 2)
	r[2] = utils.KeyGen(toUserId, 2, 1)
	r[3] = utils.KeyGen(toUserId, 2, 2)
	return r
}
