package db

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID   int64  `json:"user_id"`
	ToUserID int64  `json:"to_user_id"`
	Content  string `json:"content"`
}
