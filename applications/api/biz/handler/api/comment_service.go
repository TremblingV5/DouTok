// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/TremblingV5/DouTok/applications/api/initialize"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/hertz-contrib/jwt"

	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
)

// CommentAction .
//
//	@Tags		Comment评论
//
//	@Summary	添加或删除评论
//	@Description
//	@Param		req		body		api.DouyinCommentActionRequest	true	"评论操作信息"
//	@Success	200		{object}	comment.DouyinCommentActionResponse
//	@Failure	default	{object}	api.DouyinCommentActionResponse
//	@router		/douyin/comment/action [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildCommendActionResp(errno.ErrBind))
		return
	}

	userId := int64(jwt.ExtractClaims(ctx, c)[initialize.AuthMiddleware.IdentityKey].(float64))
	rpcReq := comment.DouyinCommentActionRequest{
		VideoId:     req.VideoId,
		ActionType:  req.ActionType,
		UserId:      userId,
		CommentId:   req.CommentId,
		CommentText: req.CommentText,
	}

	resp, err := rpc.CommentAction(ctx, rpc.CommentClient, &rpcReq)
	if err != nil {
		handler.SendResponse(c, handler.BuildCommendActionResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// CommentList .
//
//	@Tags		Comment评论
//
//	@Summary	获取某个视频之下的评论列表
//	@Description
//	@Param		req		query		api.DouyinCommentListRequest	true	"获取"
//	@Success	200		{object}	comment.DouyinCommentListResponse
//	@Failure	default	{object}	api.DouyinCommentListResponse
//	@router		/douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildCommendListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.CommentList(ctx, rpc.CommentClient, &comment.DouyinCommentListRequest{
		VideoId: req.VideoId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildCommendListResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}
