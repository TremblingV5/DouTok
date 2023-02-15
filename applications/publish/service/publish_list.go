package service

import (
	"strconv"

	"github.com/TremblingV5/DouTok/applications/publish/typedef"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/misc"
)

func QueryPublishListInHBase(user_id int64) ([]typedef.VideoInHB, error) {
	user_id_string := strconv.FormatInt(user_id, 10)
	user_id_string = misc.LFill(user_id_string, 6)

	filters := hbaseHandle.GetFilterByRowKeyPrefix(user_id_string)

	video_list, err := HBClient.Scan(
		"list", filters...,
	)

	list := []typedef.VideoInHB{}
	if err != nil {
		return list, err
	}

	for _, v := range video_list {
		temp := typedef.VideoInHB{}
		err := misc.Map2Struct4HB(v, &temp)
		if err != nil {
			continue
		}
		list = append(list, temp)
	}

	return list, nil
}
