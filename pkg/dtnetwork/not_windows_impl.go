//go:build !windows
// +build !windows

package dtnetwork

import (
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

func GetTransporter(enableNetpoll bool) func(options *config.Options) network.Transporter {
	if enableNetpoll {
		return netpoll.NewTransporter
	}
	return standard.NewTransporter
}
