package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// 在日志中记录每次请求 api gateway 的信息
func CacheAPIRequest() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		hlog.Infof("uri = %s\nheader = %s\nbody = %s\n", ctx.Request.URI().Path(), ctx.Request.Header.String(), ctx.Request.Body())
	}
}
