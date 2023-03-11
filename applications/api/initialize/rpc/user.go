package rpc

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/LogBuilder"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var UserClient userservice.Client

// User RPC 客户端初始化
func initUserRpc(Config *dtviper.Config) {
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

	c, err := userservice.NewClient(
		ServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r),                            // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(err)
	}
	UserClient = c
}

// 传递 注册操作 的上下文, 并获取 RPC Server 端的响应.
func Register(ctx context.Context, userClient userservice.Client, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp, err = userClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

// 传递 登录操作 的上下文, 并获取 RPC Server 端的响应.
func Login(ctx context.Context, userClient userservice.Client, req *user.DouyinUserLoginRequest) (int64, error) {
	log := LogBuilder.InitLogBuilder("info", "RPC request success")
	defer log.Write(logHandle)

	log.Collect("username", req.Username)

	resp, err := userClient.Login(ctx, req)
	if resp != nil {
		log.Collect("user_id", fmt.Sprint(resp.UserId))
	}

	if err != nil {
		log.SetLogType("error")
		log.SetMessage("RPC request defeat")
		log.Collect("ErrorInfo", err.Error())
		return 0, err
	}
	if resp.StatusCode != 0 {
		log.SetLogType("warn")
		log.SetMessage("RPC request success with error number except success")
		log.Collect("ErrorNo", fmt.Sprint(resp.StatusCode))
		log.Collect("ErrorMsg", resp.StatusMsg)
		return 0, errno.NewErrNo(int(resp.StatusCode), resp.StatusMsg)
	}
	return resp.UserId, nil
}

// 传递 获取用户信息操作 的上下文, 并获取 RPC Server 端的响应.
func GetUserById(ctx context.Context, userClient userservice.Client, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp, err = userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}
