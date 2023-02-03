package rpc

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var commentClient commentservice.Client

func InitCommentRpc() {
	addr := ClientConfig.Etcd.Address + ":" + ClientConfig.Etcd.Port
	registry, err := etcd.NewEtcdResolver([]string{addr})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		"comment",
		client.WithResolver(registry),
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
