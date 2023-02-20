package service

import (
	"encoding/binary"
	"strconv"
)

type VideoInHB struct {
	Id         []byte `json:"id"`
	AuthorId   []byte `json:"author_id"`
	AuthorName []byte `json:"author_name"`
	Title      []byte `json:"title"`
	VideoUrl   []byte `json:"video_url"`
	CoverUrl   []byte `json:"cover_url"`
	Timestamp  []byte `json:"timestamp"`
}

func ToInt64(data []byte) int64 {
	return int64(binary.BigEndian.Uint64(data))
}

func (v *VideoInHB) GetId() int64 {
	id_string := string(v.Id)
	i, _ := strconv.ParseInt(id_string, 10, 64)
	return i
}

func (v *VideoInHB) GetAuthorId() int64 {
	id_string := string(v.AuthorId)
	i, _ := strconv.ParseInt(id_string, 10, 64)
	return i
}

func (v *VideoInHB) GetVideoUrl() string {
	return string(v.VideoUrl)
}

func (v *VideoInHB) GetTitle() string {
	return string(v.Title)
}

func (v *VideoInHB) GetCoverUrl() string {
	return string(v.CoverUrl)
}

func (v *VideoInHB) GetTimestamp() string {
	return string(v.Timestamp)
}

func (v *VideoInHB) GetAuthorName() string {
	return string(v.AuthorName)
}
