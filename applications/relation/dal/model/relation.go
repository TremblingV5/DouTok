package model

import "time"

// 关系表
type Relation struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	ToUserId  int64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
