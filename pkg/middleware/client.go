package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

// 让编译器去检查ClientMiddleware和endpoint.Middleware是不是相同类型，如果不是则会报错
var _ endpoint.Middleware = ClientMiddleware

// ClientMiddleware client middleware print server address 、rpc timeout and connection timeout
// 相当于对Endpoint进行包装，在调用前输出一些信息
func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get server information
		klog.Infof("server address: %v, rpc timeout: %v, readwrite timeout: %v", ri.To().Address(), ri.Config().RPCTimeout(), ri.Config().ConnectTimeout())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}
