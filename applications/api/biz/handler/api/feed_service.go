// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/TremblingV5/DouTok/applications/api/initialize"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/hertz-contrib/jwt"

	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
)

// GetUserFeed .
// @router /douyin/feed [GET]
func GetUserFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildGetUserFeedResp(errno.ErrBind))
		return
	}

	userId := int64(0)
	if req.Token != "" {
		userId = int64(jwt.ExtractClaims(ctx, c)[initialize.AuthMiddleware.IdentityKey].(float64))
	}

	resp, err := rpc.GetUserFeed(ctx, rpc.FeedClient, &feed.DouyinFeedRequest{
		LatestTime: req.LatestTime,
		UserId:     userId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildGetUserFeedResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}
