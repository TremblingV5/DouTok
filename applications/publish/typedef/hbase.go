package typedef

import "encoding/binary"

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
	return ToInt64(v.Id)
}

func (v *VideoInHB) GetAuthorId() int64 {
	return ToInt64(v.AuthorId)
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
