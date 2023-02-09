package main

import (
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/applications/relation/handler"
	rpc6 "github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

var Logger = dlog.InitLog(6)

func Init() {
	conf.InitConfig("./config")
	db.InitDB()
	redis.InitRedis()
	rpc6.InitUserRpc()
	rpc6.InitMessageRpc()
}

func main() {

	etcdConf := conf.RelationConfig.GetStringMapString("etcd")
	serviceConf := conf.RelationConfig.GetStringMapString("server")
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
