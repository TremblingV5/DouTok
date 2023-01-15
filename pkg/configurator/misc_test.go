package configurator

import (
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	_, err := GetConfigPath("redis.yaml")
	if err != nil {
		panic(err)
	}
}
