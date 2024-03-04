package initHelper

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/commentDomain/commentdomainservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type CommentDomainClient struct {
	client commentdomainservice.Client
}

func InitCommentDomainRPCClient() *CommentDomainClient {
	config := dtviper.ConfigInit("DOUTOK_COMMENTDOMAIN", "commentDomain", nil)
	c, err := commentdomainservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(config)...)
	if err != nil {
		panic(err)
	}
	return &CommentDomainClient{client: c}
}

func (c *CommentDomainClient) AddComment(ctx context.Context, req *commentDomain.DoutokAddCommentReq) (*commentDomain.DoutokAddCommentResp, error) {
	resp, err := c.client.AddComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *CommentDomainClient) RmComment(ctx context.Context, req *commentDomain.DoutokRmCommentReq) (*commentDomain.DoutokAddCommentResp, error) {
	resp, err := c.client.RmComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *CommentDomainClient) ListComment(ctx context.Context, req *commentDomain.DoutokListCommentReq) (*commentDomain.DoutokListCommentResp, error) {
	resp, err := c.client.ListComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *CommentDomainClient) CountComment(ctx context.Context, req *commentDomain.DoutokCountCommentReq) (*commentDomain.DoutokCountCommentResp, error) {
	resp, err := c.client.CountComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
