package handlers

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/api/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"strconv"
)

// 传递 点赞操作 的上下文至 Favorite 服务的 RPC 客户端, 并获取相应的响应.
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var paramVar FavoriteActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.VideoId = int64(vid)
	paramVar.ActionType = int32(act)

	resp, err := rpc.FavoriteAction(ctx, &favorite.DouyinFavoriteActionRequest{
		VideoId:    paramVar.VideoId,
		Token:      paramVar.Token,
		ActionType: paramVar.ActionType,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// 传递 获取点赞列表操作 的上下文至 Favorite 服务的 RPC 客户端, 并获取相应的响应.
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.FavoriteList(ctx, &favorite.DouyinFavoriteListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
