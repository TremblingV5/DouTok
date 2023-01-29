package main

import (
	"net"

	"github.com/TremblingV5/DouTok/applications/publish/service"
	"github.com/TremblingV5/DouTok/kitex_gen/publish/publishservice"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	service.InitDb()
	service.InitOSS()
}

func main() {
	Init()

	registry, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {

	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {

	}

	svr := publishservice.NewServer(
		new(PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
	)

	if err := svr.Run(); err != nil {

	}
}
