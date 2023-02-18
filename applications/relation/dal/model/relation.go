package model

import "time"

// 关系表
type Relation struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64 `gorm:"index:user_id"`
	ToUserId  int64 `gorm:"index:to_user_id"`
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
