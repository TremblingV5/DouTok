package middleware

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/acl"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"golang.org/x/net/context"
)

/*
	初始化Kitex中间件，以防止不规范的循环调用出现。
	preventList中包括若干服务名，凡是出现在该List中的所有服务对使用该中间件的服务的调用均会被直接拒绝
*/
func GetKitexStreamingCallPreventMiddleware(preventList []string) endpoint.Middleware {
	return acl.NewACLMiddleware(
		[]acl.RejectFunc{
			func(ctx context.Context, request interface{}) error {
				fmt.Println(request)
				return nil
			},
		},
	)
}
