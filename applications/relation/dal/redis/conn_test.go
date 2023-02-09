package redis

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"testing"
)

func TestConnRedis(t *testing.T) {
	v, err := conf.InitConfig("./", "relation")
	if err != nil {
		t.Fatal(err)
	}
	Conn(v)
	_, err = RD.Get(context.Background(), "scz00")
	if err != nil {
		fmt.Println("err:", err)
	}
	//fmt.Println("result:", err)

}
