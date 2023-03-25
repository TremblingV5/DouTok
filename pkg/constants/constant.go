package constants

const (
	IdentityKey = "user_id"

	/*
		如下是数据库名的定义
	*/
	DbDefault   = "default"
	FeedSendBox = "SendBox"
	TimeCache   = "MarkedTime"

	AddrPrefix = "user_addr"

	ErrConfigFileNotFound = "Config file not found"
	// Redis 关注数字段
	FollowCount = "follow_count-"
	// Redis 粉丝数字段
	FollowerCount = "follower_count-"
	// 关注列表
	FollowListPrefix = "follow_list-"
	// 粉丝列表
	FollowerListPrefix = "follower_list-"
)
