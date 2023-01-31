package dao

import (
	"gorm.io/gorm"
	"time"
)

// 评论表
type Comment struct {
	ID        int64 `gorm:"auto_increment;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	UserID    int64
	VideoID   int64  `gorm:"index"`
	Content   string `gorm:"type:text"`
}
