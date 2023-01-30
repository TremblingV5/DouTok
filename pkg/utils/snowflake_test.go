package utils

import (
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"testing"
	"time"
)

func TestGetSnowFlakeId(t *testing.T) {
	for i := 1; i <= 3; i++ {
		go printId(int64(i * 100))
	}
	time.Sleep(1 * time.Second)
}

func TestGetSnowFlakeId2(t *testing.T) {
	config := dtviper.ConfigInit("TEST_SNOWFLAKE", "user")
	node := config.Viper.GetInt64("Snowflake.Node")
	InitSnowFlake(node)
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
}

func printId(node int64) {
	InitSnowFlake(node)
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
	fmt.Printf("%d %d\n", node, GetSnowFlakeId())
}
