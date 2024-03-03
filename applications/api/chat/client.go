package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"github.com/jellydator/ttlcache/v2"

	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

const (
// Time allowed to write a message to the peer.
// writeWait = 10 * time.Second

// Time allowed to read the next pong message from the peer.
// pongWait = 60 * time.Second

// Send pings to peer with this period. Must be less than pongWait.
// pingPeriod = (pongWait * 9) / 10

// Maximum message size allowed from peer.
// maxMessageSize = 512
)

var (
	// newline  = []byte{'\n'}
	// space    = []byte{' '}
	upgrader = websocket.HertzUpgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(ctx *app.RequestContext) bool {
			return true
		},
	}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	// The websocket connection.
	conn *websocket.Conn

	// to user id
	userId string
}

type ClientMsg struct {
	UserId     int64  `json:"user_id"`
	ToUserId   int64  `json:"to_user_id"`
	MsgContent string `json:"msg_content"`
}

type ServerMsg struct {
	FromUserId int64  `json:"from_user_id"`
	MsgContent string `json:"msg_content"`
}

// websocket 服务端的实现
// serveWs handles websocket requests from the peer.
func ServeWs(ctx context.Context, c *app.RequestContext) {

	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {

		// 注册 client

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}
			// TODO 这里消息编解码可能有问题，需要考虑到客户端的处理方式
			clientMsg := ClientMsg{}
			if err := json.Unmarshal(msg, &clientMsg); err != nil {
				continue
			}

			clientFrom, err := hub.clients.Get(fmt.Sprint(clientMsg.UserId))
			if errors.Is(err, ttlcache.ErrNotFound) {
				// 注册 client
				client := &Client{conn: conn, userId: fmt.Sprint(clientMsg.UserId)}
				hub.register <- client
			} else {
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("error: %v", err)
					}
					break
				}
			}
			// 向 message 模块发送消息
			resp, err := rpc.MessageAction(ctx, rpc.MessageClient, &message.DouyinMessageActionRequest{
				ToUserId:   clientMsg.ToUserId,
				ActionType: 1,
				Content:    clientMsg.MsgContent,
			})
			if err != nil {
				handler.SendResponse(c, handler.BuildMessageActionResp(errno.ConvertErr(err)))
				return
			}
			// 获取 B 用户的连接并发送消息
			clientTo, err := hub.clients.Get(fmt.Sprint(clientMsg.ToUserId))
			if errors.Is(err, ttlcache.ErrNotFound) {
				// B 不在线
				handler.SendResponse(c, handler.BuildMessageActionResp(errno.Success))
				return
			} else {
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						hlog.Info("error: %v", err)
					}
					break
				} else {
					// B 在线
					serverMsg := ServerMsg{
						FromUserId: clientMsg.UserId,
						MsgContent: clientMsg.MsgContent,
					}
					data, err := json.Marshal(serverMsg)
					if err != nil {
						handler.SendResponse(c, handler.BuildMessageActionResp(errno.ConvertErr(err)))
						return
					}
					if err := clientTo.(Client).conn.WriteMessage(websocket.TextMessage, data); err != nil {
						return
					}
				}
			}
			// 返回 websocket 响应
			data, err := json.Marshal(resp)
			if err != nil {
				handler.SendResponse(c, handler.BuildMessageActionResp(errno.ConvertErr(err)))
				return
			}
			if err := clientFrom.(Client).conn.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		}
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildMessageActionResp(errno.ConvertErr(err)))
	}
}

// 实现 socket 消息推送中心
func SocketChat(ctx context.Context, c *app.RequestContext) {
	/**
	conn 等价于最初 hertz-server 握手之后创建的 conn，
	如果能够维护一个map，在保证对应conn在keep-alive期间，是可以实现消息的转发操作的
	是否有必要？
	*/
	//conn := c.GetConn()
}
