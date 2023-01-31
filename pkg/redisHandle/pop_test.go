package redishandle

import (
	"context"
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
)

func TestPop(t *testing.T) {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)

	redisCaches, _ := InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)

	redisCaches[constants.DbDefault].LPush(context.Background(), "test_push", "lpush")
	redisCaches[constants.DbDefault].RPush(context.Background(), "test_push", "rpush")

	result, err := redisCaches[constants.DbDefault].LPops(context.Background(), "test_push", 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
