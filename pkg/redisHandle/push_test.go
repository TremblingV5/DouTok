package redishandle

import (
	"context"
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
)

func TestPush(t *testing.T) {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)

	redisCaches, _ := InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)

	init, _ := redisCaches[constants.DbDefault].ListSize(context.Background(), "test_push")

	if err := redisCaches[constants.DbDefault].LPush(context.Background(), "test_push", "lpush"); err != nil {
		panic(err)
	}

	if err := redisCaches[constants.DbDefault].RPush(context.Background(), "test_push", "rpush"); err != nil {
		panic(err)
	}

	res, _ := redisCaches[constants.DbDefault].ListSize(context.Background(), "test_push")

	if init+2 != res {
		panic("Defeat")
	}

	fmt.Println(init, res)
}
