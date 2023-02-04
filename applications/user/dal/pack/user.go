package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/user/db"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

// User pack user info
// db.User结构体是用于和数据库交互的
// user.User结构体是用于RPC传输信息的
func User(ctx context.Context, u *db.User, fromID int64) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "已注销用户",
		}, nil
	}

	follow_count := int64(u.FollowingCount)
	follower_count := int64(u.FollowerCount)

	// true->fromID已关注u.ID，false-fromID未关注u.ID
	isFollow := false
	//relation, err := GetRelation(ctx, fromID, int64(u.ID))
	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	return nil, err
	//}
	//
	//if relation != nil {
	//	isFollow = true
	//}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   follow_count,
		FollowerCount: follower_count,
		IsFollow:      isFollow,
	}, nil
}

// GetRelation get relation info
//func GetRelation(ctx context.Context, uid int64, tid int64) (*Relation, error) {
//	relation := new(Relation)
//
//	if err := DB.WithContext(ctx).First(&relation, "user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
//		return nil, err
//	}
//	return relation, nil
//}

// Users pack list of user info
func Users(ctx context.Context, us []*db.User, fromID int64) ([]*user.User, error) {
	users := make([]*user.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u, fromID)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
