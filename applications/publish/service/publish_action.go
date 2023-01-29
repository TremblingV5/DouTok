package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/TremblingV5/DouTok/applications/publish/dal/model"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/misc"
)

func SavePublish(user_id int64, title string, data []byte) error {
	timestamp := time.Now().Unix()

	// 1. 上传封面和视频到OSS
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprint(user_id) + title))
	filename := hex.EncodeToString(hasher.Sum(nil))

	OSSClient.Put(
		"video", filename, bytes.NewReader(data),
	)

	play_url := "https://" + OssCfg.BucketName + "." + OssCfg.Endpoint + "/video/" + filename
	cover_url := play_url + "?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto"

	// 2. 写入数据到MySQl
	id, err := SaveVideo2DB(
		uint64(user_id), title, play_url, cover_url,
	)
	if err != nil {
		return err
	}

	// 3. 写入数据到HBase，分别写入publish表和feed表

	if err := SaveVideo2HB(id, uint64(user_id), title, play_url, cover_url, fmt.Sprint(timestamp)); err != nil {
		dlog.Warn(err)
	}

	return nil
}

func SaveVideo2DB(user_id uint64, title string, play_url string, cover_url string) (uint64, error) {
	newVideo := model.Video{
		AuthorID: user_id,
		Title:    title,
		VideoUrl: play_url,
		CoverUrl: cover_url,
		FavCount: 0,
		ComCount: 0,
	}

	err := Video.Create(&newVideo)

	if err != nil {
		return 0, err
	}

	return newVideo.ID, nil
}

func SaveVideo2HB(id uint64, user_id uint64, title string, play_url string, cover_url string, timestamp string) error {
	// newVideo := typedef.VideoInHB{
	// 	Id:         int64(id),
	// 	AuthorId:   int64(user_id),
	// 	AuthorName: "",
	// 	Title:      title,
	// 	VideoUrl:   play_url,
	// 	CoverUrl:   cover_url,
	// 	Timestamp:  timestamp,
	// }

	publish_rowkey := misc.LFill(fmt.Sprint(user_id), 6) + timestamp
	feed_rowkey := timestamp + misc.LFill(fmt.Sprint(user_id), 6)

	hbData := map[string]map[string][]byte{
		"data": {
			"id":          []byte(fmt.Sprint(id)),
			"author_id":   []byte(fmt.Sprint(user_id)),
			"author_name": []byte(""),
			"title":       []byte(title),
			"video_url":   []byte(play_url),
			"cover_url":   []byte(cover_url),
		},
	}

	if err := HBClient.Put(
		"publish", publish_rowkey, hbData,
	); err != nil {

	}

	if err := HBClient.Put(
		"feed", feed_rowkey, hbData,
	); err != nil {

	}

	return nil
}
