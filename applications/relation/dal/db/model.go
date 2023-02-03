package db

import (
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Follow struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"UniqueIndex:idx_user_follow_deleteAt"`
	UserId    int64                 `gorm:"UniqueIndex:idx_user_follow_deleteAt"`
	FollowId  int64                 `gorm:"UniqueIndex:idx_user_follow_deleteAt"`
}

type Follower struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  soft_delete.DeletedAt `gorm:"UniqueIndex:idx_user_follow_deleteAt"`
	UserId     int64                 `gorm:"UniqueIndex:idx_user_follow_deleteAt"`
	FollowerId int64                 `gorm:"UniqueIndex:idx_user_follow_deleteAt"`
}

func Insert2FollowTable(userID, followID int64) error {
	return DB.Create(&Follow{
		FollowId: followID,
		UserId:   userID,
	}).Error

}

func Insert2FollowerTable(userID, followerID int64) error {
	return DB.Create(&Follower{
		FollowerId: followerID,
		UserId:     userID,
	}).Error
}

func AddFollowNum(userID int64) error {

	return DB.Model(&user.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
}

func AddFollowerNum(userID int64) error {
	return DB.Model(&user.User{}).Where("id = ?", userID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
}

func DeleteOnFollowTable(userID, followID int64) error {
	return DB.Model(&Follow{}).Where("user_id = ? and follow_id = ?", userID, followID).Delete(&Follow{}).Error
}

func DeleteOnFollowerTable(userID, followerID int64) error {
	return DB.Where("user_id = ? and follower_id = ?", userID, followerID).Delete(&Follower{}).Error
}

func DecrFollowNum(userID int64) error {
	return DB.Model(&user.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error

}
func DecrFollowerNum(userID int64) error {
	return DB.Model(&user.User{}).Where("id = ?", userID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error

}
func IsRelation(userID, toUserID int64) error {
	return DB.Where("user_id = ? and follow_id = ?", userID, toUserID).First(&Follow{}).Error
}
func Relation(userID, toUserID int64) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := IsRelation(userID, toUserID); err == nil {
		tx.Rollback()
		return errno.RelationRepeatedErr
	}
	if err := Insert2FollowTable(userID, toUserID); err != nil {
		tx.Rollback()
		return err
	}
	if err := Insert2FollowerTable(toUserID, userID); err != nil {
		tx.Rollback()
		return err
	}
	if err := AddFollowNum(userID); err != nil {
		tx.Rollback()
		return err
	}
	if err := AddFollowerNum(toUserID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func CancelRelation(userId, toUserId int64) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := IsRelation(userId, toUserId); err != nil {
		tx.Rollback()
		return errno.RelationUnfollowedErr
	}
	if err := DeleteOnFollowTable(userId, toUserId); err != nil {
		tx.Rollback()
		return err
	}
	if err := DeleteOnFollowerTable(toUserId, userId); err != nil {
		tx.Rollback()
		return err
	}
	if err := DecrFollowNum(userId); err != nil {
		tx.Rollback()
		return err
	}
	if err := DecrFollowerNum(toUserId); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetFollowList(userId int64) ([]*user.User, error) {
	var userIds []int64
	var users []*user.User
	err := DB.Model(&Follow{}).Select("follow_id").Where("user_id = ?", userId).Find(&userIds).Error
	if err != nil {
		return nil, err
	}
	if len(userIds) == 0 {
		return users, nil
	}
	err = DB.Find(&users, userIds).Error
	return users, err
}

func GetFollowerList(userId int64) ([]*user.User, error) {
	var userIds []int64
	var users []*user.User
	err := DB.Model(&Follower{}).Select("follower_id").Where("user_id = ?", userId).Find(&userIds).Error
	if err != nil {
		return nil, err
	}
	if len(userIds) == 0 {
		return users, nil
	}
	err = DB.Find(&users, userIds).Error
	return users, err
}
