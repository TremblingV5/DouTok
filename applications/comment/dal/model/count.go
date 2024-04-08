package model

import "time"

type CommentCount struct {
	Id        int64 `gorm:"index:com_count_id"` // 与Comment表的Id相同
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
