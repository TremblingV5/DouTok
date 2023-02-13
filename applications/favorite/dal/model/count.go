package model

import "time"

type FavoriteCount struct {
	VideoId   int64
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
