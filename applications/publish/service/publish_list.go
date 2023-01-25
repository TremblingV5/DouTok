package service

import (
	"strconv"

	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/misc"
)

func QueryPublishListInHBase(user_id int64) {
	user_id_string := strconv.FormatInt(user_id, 10)
	user_id_string = misc.LFill(user_id_string, 6)

	filters := hbaseHandle.GetFilterByRowKeyPrefix(user_id_string)

	video_list, err := HBClient.Scan(
		"list", filters...,
	)
}
