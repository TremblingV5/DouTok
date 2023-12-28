package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SnowflakeHandle struct {
	node *snowflake.Node
}

func NewSnowflakeHandle(node int64) *SnowflakeHandle {
	n, err := snowflake.NewNode(node)
	if err != nil {
		panic(err)
	}
	return &SnowflakeHandle{
		node: n,
	}
}

func (h *SnowflakeHandle) GetId() snowflake.ID {
	return h.node.Generate()
}

var defaultNode *snowflake.Node

// 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID
// 传入的 node 用于控制 10bit 的 NodeID，确保不同机器唯一
func InitSnowFlake(node int64) {
	n, err := snowflake.NewNode(node)
	if err != nil {
		klog.Info(err)
		return
	}
	defaultNode = n
}

func GetSnowFlakeId() snowflake.ID {
	return defaultNode.Generate()
}
