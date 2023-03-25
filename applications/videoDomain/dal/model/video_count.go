package model

import "time"

type VideoCount struct {
	Id           uint64 `gorm:"primarykey"`
	UserId       uint64
	PublishCount int64
	CreatedAt    time.Time
	UpdateAt     time.Time
}
