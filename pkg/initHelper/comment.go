package initHelper

import (
	"errors"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"golang.org/x/net/context"
)

type CommentClient struct {
	client commentservice.Client
}

func InitCommentRPCClient() *CommentClient {
	config := dtviper.ConfigInit("DOUTOK_COMMENT", "comment")
	c, err := commentservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(&config)...)
	if err != nil {
		panic(err)
	}
	return &CommentClient{client: c}
}

func (c *CommentClient) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (*comment.DouyinCommentActionResponse, error) {
	resp, err := c.client.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *CommentClient) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	resp, err := c.client.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
