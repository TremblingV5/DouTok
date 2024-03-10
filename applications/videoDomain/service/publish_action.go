package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/TremblingV5/DouTok/applications/videoDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/pkg/LogBuilder"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

type SavePublishService struct {
	ctx context.Context
}

func NewSavePublishService(ctx context.Context) *SavePublishService {
	return &SavePublishService{
		ctx: ctx,
	}
}

func (s *SavePublishService) SavePublish(user_id int64, title string, data []byte) error {
	log := LogBuilder.InitLogBuilder()
	defer log.Write(Logger)
	log.Collect("user_id", strconv.FormatInt(user_id, 10))

	timestamp := time.Now().Unix()

	// 1. 上传封面和视频到OSS
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprint(user_id) + title))
	filename := hex.EncodeToString(hasher.Sum(nil)) + ".mp4"

	if err := MinioClient.Put(
		"video", filename, bytes.NewReader(data), len(data),
	); err != nil {
		log.SetLogType("error")
		log.SetMessage("Put video to OSS failed")
		log.Collect("errMsg", err.Error())
		return err
	}

	play_url := "http://" + DomainConfig.MinIO.Endpoint + "/" + DomainConfig.MinIO.Bucket + "/doutok/video/" + filename
	cover_url := play_url + "?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto"

	// 2. 写入数据到MySQl
	id, err := SaveVideo2DB(
		uint64(user_id), title, play_url, cover_url,
	)
	if err != nil {
		log.SetLogType("error")
		log.SetMessage("Save video info to db failed")
		log.Collect("errMsg", err.Error())
		return err
	}

	// 3. 写入数据到HBase，分别写入publish表和feed表
	err = SaveVideo2HB(id, uint64(user_id), title, play_url, cover_url, fmt.Sprint(timestamp))
	if err != nil {
		return err
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

// SaveVideo2HB TODO 这里的错误error需要处理
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

	err := HBClient.Put(
		"publish", publish_rowkey, hbData,
	)
	if err != nil {
		return nil
	}
	err = HBClient.Put(
		"feed", feed_rowkey, hbData,
	)
	if err != nil {
		return nil
	}
	return nil
}
