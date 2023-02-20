package model

import "time"

type FavoriteCount struct {
	VideoId   int64 `gorm:"index:video_id"`
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
