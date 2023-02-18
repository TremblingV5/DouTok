package model

import "time"

type Favorite struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64 `gorm:"index:user_id"`
	VideoId   int64 `gorm:"index:video_id"`
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
