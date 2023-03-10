// Code generated by hertz generator.

package main

import (
	"context"
	handler "github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		c.String(consts.StatusOK, "no method")
	})
}
