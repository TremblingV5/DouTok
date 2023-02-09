package db

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/bytedance/gopkg/util/logger"
	"testing"
)

func TestSQL(t *testing.T) {
	//读取配置
	v, err := conf.InitConfig("./", "relation")
	if err != nil {
		logger.Fatal(err)
	}
	//连接数据库
	if err := Conn(v); err != nil {
		logger.Fatal(err)
	}
	r, _ := GetFollowList(1)
	fmt.Println(r)
	r2, _ := GetFollowerList(2)
	fmt.Println(r2)
}
