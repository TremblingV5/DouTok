package service

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"testing"
)

func TestQueryVideoCountFromRDB(t *testing.T) {
	misc.InitViperConfig()

	InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.HBase"),
	)

	v, err := QueryVideoFromRBDById(1)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(v)
}

func TestQueryAFewVideoCountFromRDB(t *testing.T) {
	misc.InitViperConfig()

	InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.HBase"),
	)

	videos, err := QuerySomeVideoFromRDBByIds(1, 22)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(videos)
}
