package hbaseHandle

import (
	"context"
	"io"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

func (c *HBaseClient) Scan(table string, options ...func(hrpc.Call) error) (map[string]map[string]interface{}, error) {
	req, err := hrpc.NewScanStr(
		context.Background(), table, options...,
	)

	if err != nil {
		return nil, err
	}

	scanner := c.Client.Scan(req)

	packed := make(map[string]map[string]interface{})

	for {
		res, err := scanner.Next()
		if err == io.EOF || res == nil {
			break
		}

		var rowKey string
		temp := make(map[string]interface{})

		for _, v := range res.Cells {
			rowKey = string(v.Row)
			temp[string(v.Qualifier)] = v.Value
		}

		packed[rowKey] = temp
	}

	return packed, nil
}

func (c *HBaseClient) ScanRange(table string, start string, end string, options ...func(hrpc.Call) error) (map[string]map[string]any, error) {
	req, err := hrpc.NewScanRangeStr(
		context.Background(), table, start, end, options...,
	)

	if err != nil {
		return nil, err
	}

	scanner := c.Client.Scan(req)

	packed := make(map[string]map[string]any)

	for {
		res, err := scanner.Next()
		if err == io.EOF || res == nil {
			break
		}

		var rowKey string
		temp := make(map[string]any)

		for _, v := range res.Cells {
			rowKey = string(v.Row)
			temp[string(v.Qualifier)] = string(v.Value)
		}

		packed[rowKey] = temp
	}

	return packed, nil
}

func GetFilterByRowKeyPrefix(prefix string) []func(hrpc.Call) error {
	var res []func(hrpc.Call) error

	f2 := filter.NewPrefixFilter([]byte(prefix))
	res = append(res, hrpc.Filters(f2))

	return res
}

func GetFilterByRowKeyRange(num int, start string, end string) []func(hrpc.Call) error {
	var res []func(hrpc.Call) error

	f1 := filter.NewPageFilter(int64(num))
	res = append(res, hrpc.Filters(f1))

	f2 := filter.NewRowRange([]byte(start), []byte(end), true, false)
	res = append(res, hrpc.Filters(f2))

	return res
}
