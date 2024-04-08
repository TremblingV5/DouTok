package hbaseHandle

import (
	"context"

	"github.com/tsuna/gohbase/hrpc"
)

func (c *HBaseClient) Put(table string, rowKey string, values map[string]map[string][]byte) error {
	putReq, err := hrpc.NewPutStr(
		context.Background(), table, rowKey, values,
	)

	if err != nil {
		return err
	}

	if _, err := c.Client.Put(putReq); err != nil {
		return err
	}

	return nil
}
