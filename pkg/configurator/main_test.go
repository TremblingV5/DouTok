package configurator

import (
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
)

func TestInitConfig(t *testing.T) {
	var config configStruct.RedisConfig
	InitConfig(
		&config, "redis.yaml",
	)
	fmt.Println(config)
}
