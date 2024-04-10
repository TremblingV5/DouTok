package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type RelationClient struct {
	client relationservice.Client
}

func InitRelationRPCClient() *RelationClient {
	config := dtviper.ConfigInit("DOUTOK_RELATION", "relation")
	c, err := relationservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(&config)...)
	if err != nil {
		panic(err)
	}
	return &RelationClient{client: c}
}

func (c *RelationClient) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (*relation.DouyinRelationActionResponse, error) {
	resp, err := c.client.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationClient) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (*relation.DouyinRelationFollowListResponse, error) {
	resp, err := c.client.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationClient) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (*relation.DouyinRelationFollowerListResponse, error) {
	resp, err := c.client.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationClient) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (*relation.DouyinRelationFriendListResponse, error) {
	resp, err := c.client.RelationFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
