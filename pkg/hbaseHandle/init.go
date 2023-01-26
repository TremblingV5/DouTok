package hbaseHandle

import "github.com/tsuna/gohbase"

func InitHB(host string) HBaseClient {
	c := gohbase.NewClient(host)
	return HBaseClient{
		Client: c,
	}
}
