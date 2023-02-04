package ttviper

/*
	usage:
	  - go run config_test.go --client.foo=baz
	  - TIKTOK_CLIENT_FOO=baz TIKTOK_CLIENT_ECHO=0 go run config_test.go
	  - go run config_test.go --config <path to config>
*/

import (
	"testing"
)

func TestConfigInit(t *testing.T) {
	ConfigInit("DouTok", "userConfig")
}
