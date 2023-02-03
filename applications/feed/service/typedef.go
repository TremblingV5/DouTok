package service

type VideoInHB struct {
	Id         int64  `json:"id"`
	AuthorId   int64  `json:"author_id"`
	AuthorName string `json:"author_name"`
	Title      string `json:"title"`
	VideoUrl   string `json:"video_url"`
	CoverUrl   string `json:"cover_url"`
	Timestamp  string `json:"timestamp"`
}
