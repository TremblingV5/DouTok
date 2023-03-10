// Code generated by hertz generator.

package api

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/pkg/errno"

	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
)

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishActionRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	handler.SendResponse(c, handler.BuildPublishActionResp(errno.ErrBind))
	// 	return
	// }

	req.Token = c.PostForm("token")
	req.Title = c.PostForm("title")
	fs, _ := c.FormFile("data")
	f, _ := fs.Open()
	buff := new(bytes.Buffer)
	_, err = io.Copy(buff, f)
	if err != nil {
		log.Panicln(err)
	}
	req.Data = buff.Bytes()

	// TODO 这个绑定是否能够实现二进制文件的绑定（待测试）
	resp, err := rpc.PublishAction(ctx, rpc.PublishClient, &publish.DouyinPublishActionRequest{
		Title:  req.Title,
		Data:   req.Data,
		UserId: int64(c.Keys["user_id"].(float64)),
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildPublishActionResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// PublishList .
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildPublishListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.PublishList(ctx, rpc.PublishClient, &publish.DouyinPublishListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}
