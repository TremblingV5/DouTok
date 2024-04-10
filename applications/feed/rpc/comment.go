package rpc

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/initHelper"

	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
)

var commentClient commentservice.Client

func InitCommentRpc() {
	config := dtviper.ConfigInit("DOUTOK_COMMENT", "comment")

	c, err := commentservice.NewClient(
		config.Viper.GetString("Server.Name"),
		initHelper.InitRPCClientArgs(config)...,
	)

	if err != nil {
		panic(err)
	}

	commentClient = c
}

func CommentCount(ctx context.Context, req *comment.DouyinCommentCountRequest) (resp *comment.DouyinCommentCountResponse, err error) {
	resp, err = commentClient.CommentCount(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}

	return resp, nil
}
