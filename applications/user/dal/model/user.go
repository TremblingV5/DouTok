package model

import "time"

type User struct {
	ID              uint64 `gorm:"primarykey"`
	UserName        string
	Password        string
	Salt            string
	Avatar          string
	BackgroundImage string
	Signature       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
