package model

import "time"

type User struct {
	ID            uint64    `gorm:"comment:自增主键"`
	CreateAt      time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UpdateAt      time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UserID        int64     `gorm:"type:bigint;unsigned;not null;unique;uniqueIndex:idx_user_id" json:"user_id"`
	UserName      string    `gorm:"type:varchar(50);not null;unique;uniqueIndex:idx_user_name" json:"name" validate:"min=6,max=32"`
	PassWord      string    `gorm:"type:varchar(50);not null" json:"password" validate:"min=6,max=32"`
	FollowCount   int64     `gorm:"type:bigint;unsigned;not null;default:0" json:"follow_count"`
	FollowerCount int64     `gorm:"type:bigint;unsigned;not null;default:0" json:"follower_count"`
}
