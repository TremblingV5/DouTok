package db

import (
	"context"

	"gorm.io/gorm"
)

// User Gorm Data Structures
// User和Video是多对多关系，两个模型之间会有一个连接表 user_favorite_videos
type User struct {
	gorm.Model
	UserName       string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string `gorm:"type:varchar(256);not null" json:"password"`
	FollowingCount int    `gorm:"default:0" json:"following_count"`
	FollowerCount  int    `gorm:"default:0" json:"follower_count"`
}

func (User) TableName() string {
	return "user"
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserByID
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)

	if err := DB.WithContext(ctx).First(&res, userID).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
