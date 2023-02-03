package main

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/applications/relation/handler"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func main() {

	//读取配置
	v := conf.InitConfig("./config", "relation")
	//连接数据库
	db.Conn(v)
	//连接redis
	redis.Conn(v)

	etcdConf := v.GetStringMapString("etcd")
	serviceConf := v.GetStringMapString("server")
	fmt.Println(etcdConf, serviceConf)
	r, err := etcd.NewEtcdRegistry([]string{etcdConf["host"] + ":" + etcdConf["port"]})
	addr, _ := net.ResolveTCPAddr("tcp", serviceConf["host"]+":"+serviceConf["port"])
	if err != nil {
		return
	}
	s := relationservice.NewServer(&handler.RelationServiceImpl{},
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceConf["name"]}),
		server.WithServiceAddr(addr),
		//server.WithMuxTransport(),
	)
	s.Run()

}
