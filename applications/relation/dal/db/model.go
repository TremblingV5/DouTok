package db

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/dal/model"
	"github.com/TremblingV5/DouTok/applications/relation/dal/query"
	"github.com/TremblingV5/DouTok/applications/relation/misc"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
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
	return query.FollowFollowerCount.AddFollowCount(userID)
}

func AddFollowerNum(userID int64) error {
	return query.FollowFollowerCount.AddFollowerCount(userID)
}

func DeleteOnFollowTable(userID, followID int64) error {
	return DB.Model(&Follow{}).Where("user_id = ? and follow_id = ?", userID, followID).Delete(&Follow{}).Error
}

func DeleteOnFollowerTable(userID, followerID int64) error {
	return DB.Where("user_id = ? and follower_id = ?", userID, followerID).Delete(&Follower{}).Error
}

func DecrFollowNum(userID int64) error {
	return query.FollowFollowerCount.DecrFollowCount(userID)

}
func DecrFollowerNum(userID int64) error {
	return query.FollowFollowerCount.DecrFollowerCount(userID)

}
func QUeryFollowFollowerCOunt(userId int64) (model.FollowFollowerCount, error) {
	return query.FollowFollowerCount.QueryWihtUserId(userId)
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
		return misc.RelationRepeatedErr
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
		return misc.RelationUnfollowedErr
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

func GetFollowList(ctx context.Context, userId int64) ([]*user.User, error) {
	var userIds []int64

	err := DB.Model(&Follow{}).Select("follow_id").Where("user_id = ?", userId).Find(&userIds).Error
	if err != nil {
		return nil, err
	}
	if len(userIds) == 0 {
		return nil, nil
	}
	users := make([]*user.User, len(userIds))
	for i, _ := range users {
		users[i] = &user.User{Id: userIds[i]}
		c, err := QUeryFollowFollowerCOunt(userIds[i])
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		//v.Id = userIds[i]
		users[i].FollowCount = c.FollowCount
		users[i].FollowerCount = c.FollowerCount
		users[i].IsFollow = true
		resp, err := rpc.UserClient.GetUserById(ctx, &user.DouyinUserRequest{UserId: userIds[i]})
		if err != nil {
			return users, err
		}
		users[i].Name = resp.User.Name
		users[i].Avatar = resp.User.Avatar
	}
	//todo rpc调用添加用户其他信息
	return users, err
}

func GetFollowerList(ctx context.Context, userId int64) ([]*user.User, error) {
	userIds, err := GetFollowerIds(userId)
	if err != nil {
		return nil, err
	}
	if len(userIds) == 0 {
		return nil, nil
	}
	users := make([]*user.User, len(userIds))
	for i, _ := range userIds {
		users[i] = &user.User{Id: userIds[i]}
		c, err := QUeryFollowFollowerCOunt(userIds[i])
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		//v.Id = userIds[i]
		users[i].FollowCount = c.FollowCount
		users[i].FollowerCount = c.FollowerCount
		users[i].IsFollow = IsFollow(userId, userIds[i])
		//rpc调用查询用户信息
		resp, err := rpc.UserClient.GetUserById(ctx, &user.DouyinUserRequest{UserId: userIds[i]})
		if err != nil {
			return users, err
		}
		users[i].Name = resp.User.Name
		users[i].Avatar = resp.User.Avatar
	}
	return users, err
}

func IsFollow(userId, followUserId int64) bool {
	err := DB.Model(&Follow{}).Where("user_id = ? and to_user_id = ?", userId, followUserId).First(&Follow{}).Error
	if err != nil {
		return false
	}
	return true
}

func GetFollowerIds(userId int64) ([]int64, error) {
	var userIds []int64

	err := DB.Model(&Follower{}).Select("follower_id").Where("user_id = ?", userId).Find(&userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
