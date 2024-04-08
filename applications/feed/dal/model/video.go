package model

import (
	"time"
)

type Video struct {
	ID        uint64 `gorm:"primarykey"`
	AuthorID  uint64
	Title     string
	VideoUrl  string
	CoverUrl  string
	FavCount  uint64 // 点赞数
	ComCount  uint64 // 评论数
	CreatedAt time.Time
	UpdatedAt time.Time
}
