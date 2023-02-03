package service

import (
	"strconv"

	"github.com/TremblingV5/DouTok/applications/message/dal/model"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/misc"
)

func QueryMessageChatInHBase(user_id int64, to_user_id int64) ([]model.Message, error) {
	user_id_string := strconv.FormatInt(user_id, 10)
	user_id_string = misc.LFill(user_id_string, 6)

	to_user_id_string := strconv.FormatInt(user_id, 10)
	to_user_id_string = misc.LFill(to_user_id_string, 6)

	message_rowkey := user_id_string + to_user_id_string
	filters := hbaseHandle.GetFilterByRowKeyPrefix(message_rowkey)

	message_list, err := HBClient.Scan(
		"list", filters...,
	)

	list := []model.Message{}
	if err != nil {
		return list, err
	}

	for _, v := range message_list {
		temp := model.Message{}
		err := misc.Map2Struct(v, &temp)
		if err != nil {
			continue
		}
		list = append(list, temp)
	}

	return list, nil
}
