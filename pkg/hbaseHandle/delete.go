package hbaseHandle

import (
	"context"

	"github.com/tsuna/gohbase/hrpc"
)

func (c *HBaseClient) RmByRowKey(table string, rowKey string) error {
	rmReq, err := hrpc.NewDelStr(
		context.Background(),
		table, rowKey, make(map[string]map[string][]byte),
	)

	if err != nil {
		return nil
	}

	if _, err := c.Client.Delete(rmReq); err != nil {
		return err
	}

	return nil
}
