package service

type FavReqInKafka struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
	Op      bool  `json:"op"`
}
