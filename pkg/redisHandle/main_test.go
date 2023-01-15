package redishandle

import (
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func TestInitRedis(t *testing.T) {
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
	fmt.Println(redisCaches)
}
