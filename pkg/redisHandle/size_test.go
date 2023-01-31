package redishandle

import (
	"context"
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
)

func TestSize(t *testing.T) {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)

	redisCaches, _ := InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)

	l, err := redisCaches[constants.DbDefault].ListSize(context.Background(), "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
}
