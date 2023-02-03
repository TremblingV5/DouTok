package service

import (
	"fmt"
	"time"

	"github.com/TremblingV5/DouTok/applications/message/dal/model"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/misc"
)

//保存聊天记录
func SaveMessage(user_id int64, to_user_id int64, content string) error {

	// 2. 写入数据到MySQl
	_, err := SaveMessage2DB(
		uint64(user_id), uint64(to_user_id), content, time.Now())
	if err != nil {
		return err
	}

	// 3. 写入数据到HBase，分别写入publish表和feed表

	if err := SaveMessage2HB(uint64(user_id), uint64(to_user_id), content, fmt.Sprint(time.Now())); err != nil {
		dlog.Warn(err)
	}

	return nil
}

func SaveMessage2DB(user_id uint64, to_user_id uint64, content string, timestamp time.Time) (uint64, error) {
	newMessage := model.Message{
		UserID:    user_id,
		ToUserID:  to_user_id,
		Content:   content,
		CreatedAt: timestamp,
	}

	err := Message.Create(&newMessage)

	if err != nil {
		return 0, err
	}

	return newMessage.ID, nil
}

func SaveMessage2HB(user_id uint64, to_user_id uint64, content string, timestamp string) error {

	message_rowkey := misc.LFill(fmt.Sprint(user_id), 6) + misc.LFill(fmt.Sprint(to_user_id), 6) + timestamp

	hbData := map[string]map[string][]byte{
		"data": {
			"user_id":    []byte(fmt.Sprint(user_id)),
			"to_user_id": []byte(fmt.Sprint(to_user_id)),
			"content":    []byte(content),
			"timestamp":  []byte(timestamp),
		},
	}

	if err := HBClient.Put(
		"publish", message_rowkey, hbData,
	); err != nil {

	}

	return nil
}
