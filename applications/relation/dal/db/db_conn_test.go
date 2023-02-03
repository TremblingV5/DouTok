package db

import (
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"testing"
)

func TestDBConn(t *testing.T) {
	v := conf.InitConfig("./", "relation")
	Conn(v)

}
