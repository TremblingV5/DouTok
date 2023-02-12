package model

import "time"

type CommentCount struct {
	Id        int64 // 与Comment表的Id相同
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
