package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain/relationdomainservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type RelationDomainClient struct {
	client relationdomainservice.Client
}

func InitRelationDomainRPCClient() *RelationDomainClient {
	config := dtviper.ConfigInit("DOUTOK_RELATIONDOMAIN", "relationDomain")
	c, err := relationdomainservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(config)...)
	if err != nil {
		panic(err)
	}
	return &RelationDomainClient{client: c}
}

func (c *RelationDomainClient) AddRelation(ctx context.Context, req *relationDomain.DoutokAddRelationRequest) (*relationDomain.DoutokAddRelationResponse, error) {
	resp, err := c.client.AddRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationDomainClient) ListRelation(ctx context.Context, req *relationDomain.DoutokListRelationRequest) (*relationDomain.DoutokListRelationResponse, error) {
	resp, err := c.client.ListRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationDomainClient) RmRelation(ctx context.Context, req *relationDomain.DoutokRmRelationRequest) (*relationDomain.DoutokRmRelationResponse, error) {
	resp, err := c.client.RmRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *RelationDomainClient) CountRelation(ctx context.Context, req *relationDomain.DoutokCountRelationRequest) (*relationDomain.DoutokCountRelationResponse, error) {
	resp, err := c.client.CountRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
