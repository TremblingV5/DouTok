package service

type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type UserInfo interface {
	GetUserInfo(UserId int64) User
}
