package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/message/messageservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type MessageClient struct {
	client messageservice.Client
}

func InitMessageRPCClient() *MessageClient {
	config := dtviper.ConfigInit("DOUTOK_MESSAGE", "message")
	c, err := messageservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(&config)...)
	if err != nil {
		panic(err)
	}
	return &MessageClient{client: c}
}

func (c *MessageClient) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	resp, err := c.client.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *MessageClient) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (*message.DouyinMessageActionResponse, error) {
	resp, err := c.client.MessageAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
