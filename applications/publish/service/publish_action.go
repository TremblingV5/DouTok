package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/TremblingV5/DouTok/applications/publish/dal/model"
	"github.com/TremblingV5/DouTok/applications/publish/misc"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

func SavePublish(user_id int64, title string, data []byte) error {
	timestamp := time.Now().Unix()

	// 1. 上传封面和视频到OSS
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprint(user_id) + title))
	filename := hex.EncodeToString(hasher.Sum(nil)) + ".mp4"

	if err := OSSClient.Put(
		"video", filename, bytes.NewReader(data),
	); err != nil {
		return err
	}

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
	newVideoId := utils.GetSnowFlakeId()
	newVideo := model.Video{
		ID:       uint64(newVideoId),
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

	return uint64(newVideoId), nil
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

	timestamp_int, _ := strconv.Atoi(timestamp)
	publish_rowkey := misc.FillUserId(fmt.Sprint(user_id)) + misc.GetTimeRebound(int64(timestamp_int))
	feed_rowkey := misc.GetTimeRebound(int64(timestamp_int)) + misc.FillUserId(fmt.Sprint(user_id))

	hbData := map[string]map[string][]byte{
		"data": {
			"id":          []byte(fmt.Sprint(id)),
			"author_id":   []byte(fmt.Sprint(user_id)),
			"author_name": []byte(""),
			"title":       []byte(title),
			"video_url":   []byte(play_url),
			"cover_url":   []byte(cover_url),
			"timestamp":   []byte(timestamp),
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
