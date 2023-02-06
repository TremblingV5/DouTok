package db

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserID   int64 `gorm:"uniqueIndex:idx_user_follow"`
	FollowID int64 `gorm:"uniqueIndex:idx_user_follow"`
}

type Follower struct {
	gorm.Model
	UserId     int64 `gorm:"uniqueIndex:idx_user_follower"`
	FollowerID int64 `gorm:"uniqueIndex:idx_user_follower"`
}

func Insert2FollowTable(userID, followID int64) error {
	return DB.Create(&Follow{
		FollowID: followID,
		UserID:   userID,
	}).Error

}

func Insert2FollowerTable(userID, followerID int64) error {
	return DB.Create(&Follow{
		FollowID: followerID,
		UserID:   userID,
	}).Error
}
