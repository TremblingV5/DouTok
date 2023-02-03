package redis

import (
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"testing"
)

func TestConnRedis(t *testing.T) {
	v := conf.InitConfig("./", "relation")
	Conn(v)
}
