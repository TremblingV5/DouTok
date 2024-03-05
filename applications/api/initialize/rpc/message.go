package rpc

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/message/messageservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/remote/trans/gonet"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"runtime"
	"time"
)

var MessageClient messageservice.Client

// Comment RPC 客户端初始化
func initMessageRpc(Config *dtviper.Config) {
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := Config.Viper.GetString("Server.Name")

	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(ServiceName),
	//	provider.WithExportEndpoint("localhost:4317"),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())
	options := []client.Option{
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithRPCTimeout(30 * time.Second),             // rpc timeout
		client.WithConnectTimeout(30000 * time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),   // retry
		client.WithSuite(tracing.NewClientSuite()),          // tracer
		client.WithResolver(r),                              // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	}
	if runtime.GOOS == "windows" {
		options = append(options, client.WithTransHandlerFactory(gonet.NewCliTransHandlerFactory()))
	} else {
		options = append(options, client.WithMuxConnection(1)) // mux
	}
	c, err := messageservice.NewClient(
		ServiceName,
		options...,
	)
	if err != nil {
		panic(err)
	}
	MessageClient = c
}

// 传递 发布视频操作 的上下文, 并获取 RPC Server 端的响应.
func MessageAction(ctx context.Context, messageClient messageservice.Client, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	resp, err = messageClient.MessageAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.New(int(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

// 传递 获取用户发布视频列表操作 的上下文, 并获取 RPC Server 端的响应.
func MessageChat(ctx context.Context, messageClient messageservice.Client, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	resp, err = messageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.New(int(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}
