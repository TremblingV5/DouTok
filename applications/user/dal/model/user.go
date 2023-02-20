package model

import (
	"time"
)

type User struct {
	ID              uint64    `gorm:"primarykey"`
	UserName        string    `gorm:"index:username;type:varchar(64)"`
	Password        string    `gorm:"type:varchar(64)"`
	Salt            string    `gorm:"type:varchar(64)"`
	Avatar          string    `gorm:"type:varchar(256)"`
	BackgroundImage string    `gorm:"type:varchar(256)"`
	Signature       string    `gorm:"type:varchar(512)"`
	CreatedAt       time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
