package model

import "time"

// 关注数
type FollowCount struct {
	UserId    int64 `gorm:"index:user_id"`
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 粉丝数
type FollowerCount struct {
	UserId    int64 `gorm:"index:user_id"`
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
