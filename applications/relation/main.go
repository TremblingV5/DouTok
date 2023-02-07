package main

import (
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/applications/relation/handler"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

var Logger = dlog.InitLog(6)

func main() {

	//读取配置
	v, err := conf.InitConfig("./config", "relation")
	if err != nil {
		Logger.Fatal(err)
	}
	//连接数据库
	if err := db.Conn(v); err != nil {
		Logger.Fatal(err)
	}
	//连接redis
	if err := redis.Conn(v); err != nil {
		Logger.Fatal(err)
	}

	etcdConf := v.GetStringMapString("etcd")
	serviceConf := v.GetStringMapString("server")
	logger.Info(serviceConf)
	r, err := etcd.NewEtcdRegistry([]string{etcdConf["host"] + ":" + etcdConf["port"]})
	addr, _ := net.ResolveTCPAddr("tcp", serviceConf["host"]+":"+serviceConf["port"])
	if err != nil {
		Logger.Fatal(err)
		return
	}
	s := relationservice.NewServer(&handler.RelationServiceImpl{},
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceConf["name"]}),
		server.WithServiceAddr(addr),
		//server.WithMuxTransport(),
	)
	if err := s.Run(); err != nil {
		Logger.Fatal(err)
	}

}
