package model

import (
	"time"
)

type Message struct {
	// gorm.Model
	ID        uint64    `json:"id" gorm:"primarykey"`
	UserID    uint64    `json:"from_user_id"`
	ToUserID  uint64    `json:"to_user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"create_time"`
}
