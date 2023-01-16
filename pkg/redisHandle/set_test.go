package redishandle

import (
	"context"
	"testing"
	"time"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
)

func TestSet(t *testing.T) {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)

	redisCaches, err := InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)

	if err != nil {
		panic(err)
	}

	if err := redisCaches[constants.DbDefault].Set(context.Background(), "test_key_0001", "test_value_0001", 60*time.Second); err != nil {
		panic(err)
	}
}
