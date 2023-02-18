package model

import (
	"time"
)

type Video struct {
	ID        uint64 `gorm:"primarykey"`
	AuthorID  uint64 `gorm:"index:author_id"`
	Title     string
	VideoUrl  string
	CoverUrl  string
	FavCount  uint64    // 点赞数
	ComCount  uint64    // 评论数
	CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
